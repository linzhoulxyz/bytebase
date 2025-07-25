// @generated by protoc-gen-es v2.5.2
// @generated from file v1/changelist_service.proto (package bytebase.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv2";
import type { Message } from "@bufbuild/protobuf";
import type { EmptySchema, FieldMask, Timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file v1/changelist_service.proto.
 */
export declare const file_v1_changelist_service: GenFile;

/**
 * @generated from message bytebase.v1.CreateChangelistRequest
 */
export declare type CreateChangelistRequest = Message<"bytebase.v1.CreateChangelistRequest"> & {
  /**
   * The parent resource where this changelist will be created.
   * Format: projects/{project}
   *
   * @generated from field: string parent = 1;
   */
  parent: string;

  /**
   * The changelist to create.
   *
   * @generated from field: bytebase.v1.Changelist changelist = 2;
   */
  changelist?: Changelist;

  /**
   * The ID to use for the changelist, which will become the final component of
   * the changelist's resource name.
   *
   * This value should be 4-63 characters, and valid characters
   * are /[a-z][0-9]-/.
   *
   * @generated from field: string changelist_id = 3;
   */
  changelistId: string;
};

/**
 * Describes the message bytebase.v1.CreateChangelistRequest.
 * Use `create(CreateChangelistRequestSchema)` to create a new message.
 */
export declare const CreateChangelistRequestSchema: GenMessage<CreateChangelistRequest>;

/**
 * @generated from message bytebase.v1.GetChangelistRequest
 */
export declare type GetChangelistRequest = Message<"bytebase.v1.GetChangelistRequest"> & {
  /**
   * The name of the changelist to retrieve.
   * Format: projects/{project}/changelists/{changelist}
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message bytebase.v1.GetChangelistRequest.
 * Use `create(GetChangelistRequestSchema)` to create a new message.
 */
export declare const GetChangelistRequestSchema: GenMessage<GetChangelistRequest>;

/**
 * @generated from message bytebase.v1.ListChangelistsRequest
 */
export declare type ListChangelistsRequest = Message<"bytebase.v1.ListChangelistsRequest"> & {
  /**
   * The parent, which owns this collection of changelists.
   * Format: projects/{project}
   *
   * @generated from field: string parent = 1;
   */
  parent: string;

  /**
   * Not used.
   * The maximum number of databases to return. The service may return fewer than
   * this value.
   * If unspecified, at most 50 databases will be returned.
   * The maximum value is 1000; values above 1000 will be coerced to 1000.
   *
   * @generated from field: int32 page_size = 2;
   */
  pageSize: number;

  /**
   * Not used.
   * A page token, received from a previous `ListDatabases` call.
   * Provide this to retrieve the subsequent page.
   *
   * When paginating, all other parameters provided to `ListDatabases` must match
   * the call that provided the page token.
   *
   * @generated from field: string page_token = 3;
   */
  pageToken: string;
};

/**
 * Describes the message bytebase.v1.ListChangelistsRequest.
 * Use `create(ListChangelistsRequestSchema)` to create a new message.
 */
export declare const ListChangelistsRequestSchema: GenMessage<ListChangelistsRequest>;

/**
 * @generated from message bytebase.v1.ListChangelistsResponse
 */
export declare type ListChangelistsResponse = Message<"bytebase.v1.ListChangelistsResponse"> & {
  /**
   * The changelists from the specified request.
   *
   * @generated from field: repeated bytebase.v1.Changelist changelists = 1;
   */
  changelists: Changelist[];

  /**
   * A token, which can be sent as `page_token` to retrieve the next page.
   * If this field is omitted, there are no subsequent pages.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message bytebase.v1.ListChangelistsResponse.
 * Use `create(ListChangelistsResponseSchema)` to create a new message.
 */
export declare const ListChangelistsResponseSchema: GenMessage<ListChangelistsResponse>;

/**
 * @generated from message bytebase.v1.UpdateChangelistRequest
 */
export declare type UpdateChangelistRequest = Message<"bytebase.v1.UpdateChangelistRequest"> & {
  /**
   * The changelist to update.
   *
   * The changelist's `name` field is used to identify the changelist to update.
   * Format: projects/{project}/changelists/{changelist}
   *
   * @generated from field: bytebase.v1.Changelist changelist = 1;
   */
  changelist?: Changelist;

  /**
   * The list of fields to be updated.
   *
   * @generated from field: google.protobuf.FieldMask update_mask = 2;
   */
  updateMask?: FieldMask;
};

/**
 * Describes the message bytebase.v1.UpdateChangelistRequest.
 * Use `create(UpdateChangelistRequestSchema)` to create a new message.
 */
export declare const UpdateChangelistRequestSchema: GenMessage<UpdateChangelistRequest>;

/**
 * @generated from message bytebase.v1.DeleteChangelistRequest
 */
export declare type DeleteChangelistRequest = Message<"bytebase.v1.DeleteChangelistRequest"> & {
  /**
   * The name of the changelist to delete.
   * Format: projects/{project}/changelists/{changelist}
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message bytebase.v1.DeleteChangelistRequest.
 * Use `create(DeleteChangelistRequestSchema)` to create a new message.
 */
export declare const DeleteChangelistRequestSchema: GenMessage<DeleteChangelistRequest>;

/**
 * @generated from message bytebase.v1.Changelist
 */
export declare type Changelist = Message<"bytebase.v1.Changelist"> & {
  /**
   * The name of the changelist resource.
   * Canonical parent is project.
   * Format: projects/{project}/changelists/{changelist}
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * The creator of the changelist.
   * Format: users/{email}
   *
   * @generated from field: string creator = 3;
   */
  creator: string;

  /**
   * The last update time of the changelist.
   *
   * @generated from field: google.protobuf.Timestamp update_time = 6;
   */
  updateTime?: Timestamp;

  /**
   * @generated from field: repeated bytebase.v1.Changelist.Change changes = 7;
   */
  changes: Changelist_Change[];
};

/**
 * Describes the message bytebase.v1.Changelist.
 * Use `create(ChangelistSchema)` to create a new message.
 */
export declare const ChangelistSchema: GenMessage<Changelist>;

/**
 * @generated from message bytebase.v1.Changelist.Change
 */
export declare type Changelist_Change = Message<"bytebase.v1.Changelist.Change"> & {
  /**
   * The name of a sheet.
   *
   * @generated from field: string sheet = 1;
   */
  sheet: string;

  /**
   * The source of origin.
   * 1) changelog: instances/{instance}/databases/{database}/changelogs/{changelog}.
   * 2) raw SQL if empty.
   *
   * @generated from field: string source = 2;
   */
  source: string;
};

/**
 * Describes the message bytebase.v1.Changelist.Change.
 * Use `create(Changelist_ChangeSchema)` to create a new message.
 */
export declare const Changelist_ChangeSchema: GenMessage<Changelist_Change>;

/**
 * @generated from service bytebase.v1.ChangelistService
 */
export declare const ChangelistService: GenService<{
  /**
   * Permissions required: bb.changelists.create
   *
   * @generated from rpc bytebase.v1.ChangelistService.CreateChangelist
   */
  createChangelist: {
    methodKind: "unary";
    input: typeof CreateChangelistRequestSchema;
    output: typeof ChangelistSchema;
  },
  /**
   * Permissions required: bb.changelists.get
   *
   * @generated from rpc bytebase.v1.ChangelistService.GetChangelist
   */
  getChangelist: {
    methodKind: "unary";
    input: typeof GetChangelistRequestSchema;
    output: typeof ChangelistSchema;
  },
  /**
   * Permissions required: bb.changelists.list
   *
   * @generated from rpc bytebase.v1.ChangelistService.ListChangelists
   */
  listChangelists: {
    methodKind: "unary";
    input: typeof ListChangelistsRequestSchema;
    output: typeof ListChangelistsResponseSchema;
  },
  /**
   * Permissions required: bb.changelists.update
   *
   * @generated from rpc bytebase.v1.ChangelistService.UpdateChangelist
   */
  updateChangelist: {
    methodKind: "unary";
    input: typeof UpdateChangelistRequestSchema;
    output: typeof ChangelistSchema;
  },
  /**
   * Permissions required: bb.changelists.delete
   *
   * @generated from rpc bytebase.v1.ChangelistService.DeleteChangelist
   */
  deleteChangelist: {
    methodKind: "unary";
    input: typeof DeleteChangelistRequestSchema;
    output: typeof EmptySchema;
  },
}>;

