<script setup lang="ts">
import type {
  AdminNavigationItem,
  AdminSearchGroup,
  AdminShellMessages,
} from "@yueli/ui/admin";

const route = useRoute();
const { brand: siteBrand } = useSiteRuntime();
const sidebarOpen = ref(false);
const { can, isAdministrator } = useMe();
const isDocumentEditor = computed(() => route.path.startsWith("/manage/docs/"));

const currentLabel = computed(() => {
  if (route.path === "/manage") return "控制台";
  if (route.path.startsWith("/manage/docs")) return "文档";
  if (route.path.startsWith("/manage/collections")) return "文档集";
  if (route.path.startsWith("/manage/import")) return "批量导入";
  if (route.path.startsWith("/manage/home")) return "站点设置";
  if (route.path.startsWith("/manage/assets")) return "资源策略";
  if (route.path.startsWith("/manage/comments")) return "评论";
  if (route.path.startsWith("/manage/authorization")) return "权限与申请";
  return "控制台";
});

const messages: AdminShellMessages = {
  skipToContent: "跳到主要内容",
  search: "搜索控制台",
  searchPlaceholder: "搜索页面与常用操作",
  currentLocation: "当前位置",
};

function isActive(path: string, exact = false) {
  return exact ? route.path === path : route.path.startsWith(path);
}

function closeSidebar() {
  sidebarOpen.value = false;
}

const navigation = computed<readonly AdminNavigationItem[]>(() => [
  {
    label: "控制台",
    icon: "i-tabler-dashboard",
    to: "/manage",
    active: isActive("/manage", true),
    onSelect: closeSidebar,
  },
  ...(can("docs.document.read")
    ? [
        {
          label: "文档",
          icon: "i-tabler-file-text",
          to: "/manage/docs",
          active: isActive("/manage/docs"),
          onSelect: closeSidebar,
        },
      ]
    : []),
  ...(isAdministrator.value
    ? [
        {
          label: "评论",
          icon: "i-tabler-messages",
          to: "/manage/comments",
          active: isActive("/manage/comments"),
          onSelect: closeSidebar,
        },
      ]
    : []),
  ...(can("docs.collection.manage")
    ? [
        {
          label: "文档集",
          icon: "i-tabler-stack-2",
          to: "/manage/collections",
          active: isActive("/manage/collections"),
          onSelect: closeSidebar,
        },
      ]
    : []),
  ...(can("docs.import.manage")
    ? [
        {
          label: "批量导入",
          icon: "i-tabler-file-import",
          to: "/manage/import",
          active: isActive("/manage/import"),
          onSelect: closeSidebar,
        },
      ]
    : []),
  ...(can("docs.site_settings.manage")
    ? [
        {
          label: "站点设置",
          icon: "i-tabler-settings",
          to: "/manage/home",
          active: isActive("/manage/home"),
          onSelect: closeSidebar,
        },
      ]
    : []),
  ...(can("docs.asset_settings.manage")
    ? [
        {
          label: "资源策略",
          icon: "i-tabler-database-cog",
          to: "/manage/assets",
          active: isActive("/manage/assets"),
          onSelect: closeSidebar,
        },
      ]
    : []),
  ...(isAdministrator.value
    ? [
        {
          label: "权限与申请",
          icon: "i-tabler-shield-lock",
          to: "/manage/authorization",
          active: isActive("/manage/authorization"),
          onSelect: closeSidebar,
        },
      ]
    : []),
]);

const searchGroups = computed<readonly AdminSearchGroup[]>(() => {
  const groups: AdminSearchGroup[] = [
    {
      id: "manage-pages",
      label: "管理页面",
      items: [
        {
          id: "dashboard",
          label: "控制台",
          icon: "i-tabler-dashboard",
          to: "/manage",
        },
        ...(can("docs.document.read")
          ? [
              {
                id: "documents",
                label: "管理文档",
                icon: "i-tabler-file-text",
                to: "/manage/docs",
              },
            ]
          : []),
        ...(isAdministrator.value
          ? [
              {
                id: "comments",
                label: "评论",
                icon: "i-tabler-messages",
                to: "/manage/comments",
              },
            ]
          : []),
        ...(can("docs.collection.manage")
          ? [
              {
                id: "collections",
                label: "管理文档集",
                icon: "i-tabler-stack-2",
                to: "/manage/collections",
              },
            ]
          : []),
        ...(can("docs.import.manage")
          ? [
              {
                id: "import",
                label: "批量导入",
                icon: "i-tabler-file-import",
                to: "/manage/import",
              },
            ]
          : []),
        ...(can("docs.site_settings.manage")
          ? [
              {
                id: "settings",
                label: "站点设置",
                icon: "i-tabler-settings",
                to: "/manage/home",
              },
            ]
          : []),
        ...(can("docs.asset_settings.manage")
          ? [
              {
                id: "assets",
                label: "资源策略",
                icon: "i-tabler-database-cog",
                to: "/manage/assets",
              },
            ]
          : []),
        ...(isAdministrator.value
          ? [
              {
                id: "authorization",
                label: "权限与申请",
                icon: "i-tabler-shield-lock",
                to: "/manage/authorization",
              },
            ]
          : []),
      ],
    },
  ];
  const actions = [
    ...(can("docs.document.create")
      ? [
          {
            id: "new-document",
            label: "新建文档",
            icon: "i-tabler-file-plus",
            to: "/manage/docs/new",
          },
        ]
      : []),
    ...(can("docs.import.manage")
      ? [
          {
            id: "new-import",
            label: "开始批量导入",
            icon: "i-tabler-cloud-upload",
            to: "/manage/import",
          },
        ]
      : []),
  ];
  if (actions.length) {
    groups.push({
      id: "manage-actions",
      label: "常用操作",
      items: actions,
    });
  }
  return groups;
});
</script>

<template>
  <YAdminConsoleLayout
    :navigation="navigation"
    :search-groups="searchGroups"
    :messages="messages"
    storage-key="docs-manage"
    main-id="manage-main"
    :brand-label="siteBrand"
    brand-icon="i-tabler-book"
    brand-to="/"
    :context-label="siteBrand"
    :current-label="currentLabel"
    :immersive="isDocumentEditor"
    back-to-top-label="返回顶部"
    data-docs-manage-shell
  >
    <template #account="{ collapsed }">
      <ConsumerManageAccountControl
        home-to=""
        show-appearance
        :trigger-mode="collapsed ? 'collapsed' : 'sidebar'"
      />
    </template>
    <slot />
  </YAdminConsoleLayout>
</template>

<style scoped>
@media (max-width: 640px) {
  [data-docs-manage-shell] :deep(button),
  [data-docs-manage-shell] :deep(a[href]),
  [data-docs-manage-shell] :deep(summary) {
    min-height: 44px;
  }

  [data-docs-manage-shell] :deep(button[aria-label]),
  [data-docs-manage-shell] :deep(a[aria-label]) {
    min-width: 44px;
  }
}
</style>
