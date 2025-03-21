import { v1 as uuidv1 } from "uuid";
import { computed, ref, type Ref } from "vue";
import { useEmitteryEventListener } from "@/composables/useEmitteryEventListener";
import type { CoreTabContext, TabContext } from "../types";
import type { SchemaEditorEvents } from "./index";

export type TabsContext = {
  tabMap: Ref<Map<string, TabContext>>;
  tabList: Ref<TabContext[]>;
  currentTabId: Ref<string>;
  currentTab: Ref<TabContext | undefined>;
  addTab: (coreTab: CoreTabContext, setAsCurrentTab?: boolean) => void;
  setCurrentTab: (id: string) => void;
  closeTab: (id: string) => void;
  findTab: (target: CoreTabContext) => TabContext | undefined;
};

export const useTabs = (events: SchemaEditorEvents): TabsContext => {
  const tabMap = ref(new Map<string, TabContext>());
  const tabList = computed(() => {
    return Array.from(tabMap.value.values());
  });
  const currentTabId = ref<string>("");
  const currentTab = computed(() => {
    if (!currentTabId.value) return undefined;
    return tabMap.value.get(currentTabId.value);
  });

  const addTab = (coreTab: CoreTabContext, setAsCurrentTab = true) => {
    const existedTab = findTab(coreTab);
    if (existedTab) {
      Object.assign(existedTab, coreTab);
      if (setAsCurrentTab) {
        currentTabId.value = existedTab.id;
      }
      return;
    }

    const id = uuidv1();
    const tab: TabContext = {
      id,
      ...coreTab,
    };

    tabMap.value.set(id, tab);
    if (setAsCurrentTab) {
      // It's weird that nextTick doesn't work here
      // So we need to delay a little bit longer
      requestAnimationFrame(() => {
        currentTabId.value = id;
      });
    }
  };

  const setCurrentTab = (id: string) => {
    if (tabMap.value.has(id)) {
      currentTabId.value = id;
    } else {
      currentTabId.value = "";
    }
  };

  const closeTab = (id: string) => {
    const index = tabList.value.findIndex((tab) => tab.id === id);
    if (currentTabId.value === id) {
      // Find next tab for showing.
      let nextIndex = -1;
      if (index === 0) {
        nextIndex = 1;
      } else {
        nextIndex = index - 1;
      }
      const nextTab = tabList.value[nextIndex];
      if (nextTab) {
        setCurrentTab(nextTab.id);
      } else {
        setCurrentTab("");
      }
    }
    tabMap.value.delete(id);
  };
  const findTab = (target: CoreTabContext) => {
    return tabList.value.find((tab) => {
      if (tab.type !== target.type) return false;
      if (tab.database.name !== target.database.name) return false;
      if (tab.type === "database" && target.type === "database") {
        return tab.metadata.database.name === target.metadata.database.name;
      }
      if (tab.type === "table" && target.type === "table") {
        return (
          tab.metadata.schema.name === target.metadata.schema.name &&
          tab.metadata.table.name === target.metadata.table.name
        );
      }
      if (tab.type === "view" && target.type === "view") {
        return (
          tab.metadata.schema.name === target.metadata.schema.name &&
          tab.metadata.view.name === target.metadata.view.name
        );
      }
      if (tab.type === "procedure" && target.type === "procedure") {
        return (
          tab.metadata.schema.name === target.metadata.schema.name &&
          tab.metadata.procedure.name === target.metadata.procedure.name
        );
      }
      if (tab.type === "function" && target.type === "function") {
        return (
          tab.metadata.schema.name === target.metadata.schema.name &&
          tab.metadata.function.name === target.metadata.function.name
        );
      }
      return false;
    });
  };
  useEmitteryEventListener(events, "clear-tabs", () => {
    tabMap.value.clear();
    currentTabId.value = "";
  });

  return {
    tabMap,
    tabList,
    currentTabId,
    currentTab,
    addTab,
    setCurrentTab,
    closeTab,
    findTab,
  };
};
