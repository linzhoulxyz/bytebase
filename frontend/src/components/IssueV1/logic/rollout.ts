import { planCheckRunSummaryForCheckRunList } from "@/components/PlanCheckRun/common";
import { t } from "@/plugins/i18n";
import { useCurrentUserV1, extractUserId } from "@/store";
import type { ComposedIssue } from "@/types";
import { IssueStatus, Issue_Type } from "@/types/proto/v1/issue_service";
import type { PlanCheckRun } from "@/types/proto/v1/plan_service";
import type { Task } from "@/types/proto/v1/rollout_service";
import {
  Task_Status,
  task_StatusToJSON,
  Task_Type,
} from "@/types/proto/v1/rollout_service";
import { extractDatabaseGroupName, hasProjectPermissionV2 } from "@/utils";
import { specForTask, projectOfIssue } from ".";

export const isGroupingChangeTaskV1 = (issue: ComposedIssue, task: Task) => {
  const spec = specForTask(issue.planEntity, task);
  if (!spec) {
    return false;
  }
  const databaseGroup = extractDatabaseGroupName(
    spec.changeDatabaseConfig?.target ?? ""
  );
  return databaseGroup !== "";
};

export const allowUserToEditStatementForTask = (
  issue: ComposedIssue,
  task: Task,
  planCheckRuns: PlanCheckRun[]
): string[] => {
  const user = useCurrentUserV1();
  const denyReasons: string[] = [];

  if (
    issue.type !== Issue_Type.DATABASE_CHANGE &&
    issue.type !== Issue_Type.DATABASE_DATA_EXPORT
  ) {
    denyReasons.push("Only database related issue type can be changed");
  }
  if (issue.status !== IssueStatus.OPEN) {
    denyReasons.push("The issue is not open");
  }
  const project = projectOfIssue(issue);
  if (!project.allowModifyStatement) {
    denyReasons.push("Cannot edit statement after issue created");
  }

  if (task.type === Task_Type.DATABASE_CREATE) {
    // For standard mode projects, we are not allowed to edit the database
    // creation SQL statement.
    denyReasons.push("Cannot edit database creation statement");
  }

  // if not creating, we are allowed to edit sql statement only when:
  // - user is the creator
  // - OR user has plans.update permission in the project
  denyReasons.push(...isTaskEditable(task, planCheckRuns));

  if (extractUserId(issue.creator) !== user.value.email) {
    const project = projectOfIssue(issue);
    if (!hasProjectPermissionV2(project, "bb.plans.update")) {
      denyReasons.push(
        t("issue.error.you-don-have-privilege-to-edit-this-issue")
      );
    }
  }
  return denyReasons;
};

export const isTaskEditable = (
  task: Task,
  planCheckRuns: PlanCheckRun[]
): string[] => {
  if (
    task.status === Task_Status.NOT_STARTED ||
    task.status === Task_Status.FAILED ||
    task.status === Task_Status.CANCELED
  ) {
    return [];
  }
  if (task.status === Task_Status.PENDING) {
    // If a task's status is "PENDING", its statement is editable if there
    // are at least ONE ERROR task checks.
    // Since once all its task checks are fulfilled, it might be queued by
    // the scheduler.
    // Editing a queued task's SQL statement is dangerous with kinds of race
    // condition risks.
    const summary = planCheckRunSummaryForCheckRunList(planCheckRuns);
    if (summary.errorCount > 0) {
      return [];
    }
    return [`Task is pending to run`];
  }

  return [`${task_StatusToJSON(task.status)} task is not editable`];
};

export const isTaskFinished = (task: Task): boolean => {
  return [Task_Status.DONE, Task_Status.SKIPPED].includes(task.status);
};

export const semanticTaskType = (type: Task_Type) => {
  switch (type) {
    case Task_Type.DATABASE_CREATE:
      return t("db.create");
    case Task_Type.DATABASE_DATA_UPDATE:
      return "DML";
    case Task_Type.DATABASE_SCHEMA_UPDATE:
    case Task_Type.DATABASE_SCHEMA_UPDATE_SDL:
      return "DDL";
    case Task_Type.DATABASE_SCHEMA_UPDATE_GHOST:
      return "gh-ost";
  }
  return "";
};
