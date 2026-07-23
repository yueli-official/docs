<script setup lang="ts">
import type { DropdownMenuItem } from "@nuxt/ui";
import type {
  AdminNavigationItem,
  AdminSearchGroup,
  AdminShellMessages,
} from "@yueli/ui/admin";

const route = useRoute();
const { brand: siteBrand } = useSiteRuntime();
const sidebarOpen = ref(false);
const { can, isAdministrator } = useMe();

const messages: AdminShellMessages = {
  skipToContent: "跳到主要内容",
  search: "搜索后台",
  searchPlaceholder: "搜索页面与常用操作",
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
  ...(can("docs.document.read") ? [{
    label: "文档",
    icon: "i-tabler-file-text",
    to: "/manage/docs",
    active: isActive("/manage/docs"),
    onSelect: closeSidebar,
  }] : []),
  ...(can("docs.collection.manage") ? [{
    label: "文档集",
    icon: "i-tabler-stack-2",
    to: "/manage/collections",
    active: isActive("/manage/collections"),
    onSelect: closeSidebar,
  }] : []),
  ...(can("docs.import.manage") ? [{
    label: "批量导入",
    icon: "i-tabler-file-import",
    to: "/manage/import",
    active: isActive("/manage/import"),
    onSelect: closeSidebar,
  }] : []),
  ...(can("docs.site_settings.manage") ? [{
    label: "站点设置",
    icon: "i-tabler-settings",
    to: "/manage/home",
    active: isActive("/manage/home"),
    onSelect: closeSidebar,
    children: [
      {
        label: "首页",
        icon: "i-tabler-home-cog",
        to: "/manage/home?section=home",
        onSelect: closeSidebar,
        active:
          isActive("/manage/home", true) &&
          (route.query.section ?? "home") === "home",
      },
      {
        label: "页脚",
        icon: "i-tabler-layout-bottombar",
        to: "/manage/home?section=footer",
        onSelect: closeSidebar,
        active:
          isActive("/manage/home", true) && route.query.section === "footer",
      },
      {
        label: "基础",
        icon: "i-tabler-adjustments-horizontal",
        to: "/manage/home?section=site",
        onSelect: closeSidebar,
        active:
          isActive("/manage/home", true) && route.query.section === "site",
      },
    ],
  }] : []),
  ...(can("docs.asset_settings.manage") ? [{
    label: "资源配置",
    icon: "i-tabler-database-cog",
    to: "/manage/assets",
    active: isActive("/manage/assets"),
    onSelect: closeSidebar,
  }] : []),
  ...(isAdministrator.value ? [{
    label: "权限与申请",
    icon: "i-tabler-shield-lock",
    to: "/manage/authorization",
    active: isActive("/manage/authorization"),
    onSelect: closeSidebar,
  }] : []),
]);

const workspaceMenuItems = computed<DropdownMenuItem[][]>(() => [
  [
    {
      type: "label",
      label: siteBrand.value,
    },
  ],
  [
    {
      label: "内容管理",
      icon: "i-tabler-layout-dashboard",
      type: "checkbox",
      checked: true,
      onSelect: (event: Event) => event.preventDefault(),
    },
  ],
  [
    {
      label: "打开文档站",
      icon: "i-tabler-external-link",
      to: "/",
      onSelect: closeSidebar,
    },
  ],
]);

const searchGroups: readonly AdminSearchGroup[] = [
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
      {
        id: "documents",
        label: "管理文档",
        icon: "i-tabler-file-text",
        to: "/manage/docs",
      },
      {
        id: "collections",
        label: "管理文档集",
        icon: "i-tabler-stack-2",
        to: "/manage/collections",
      },
      {
        id: "import",
        label: "批量导入",
        icon: "i-tabler-file-import",
        to: "/manage/import",
      },
      {
        id: "settings",
        label: "站点设置",
        icon: "i-tabler-settings",
        to: "/manage/home",
      },
      {
        id: "assets",
        label: "资源配置",
        icon: "i-tabler-database-cog",
        to: "/manage/assets",
      },
    ],
  },
  {
    id: "manage-actions",
    label: "常用操作",
    items: [
      {
        id: "new-document",
        label: "新建文档",
        icon: "i-tabler-file-plus",
        to: "/manage/docs/new",
      },
      {
        id: "new-import",
        label: "开始批量导入",
        icon: "i-tabler-cloud-upload",
        to: "/manage/import",
      },
    ],
  },
];
</script>

<template>
  <YAdminShell
    v-model:open="sidebarOpen"
    :navigation="navigation"
    :search-groups="searchGroups"
    :messages="messages"
    storage-key="docs-manage"
    main-id="manage-main"
    :default-size="16"
    :min-size="14"
    :max-size="20"
  >
    <template #brand="{ collapsed }">
      <UDropdownMenu
        :items="workspaceMenuItems"
        :content="{ align: 'center', collisionPadding: 12 }"
        :ui="{
          content: collapsed
            ? 'w-56'
            : 'w-(--reka-dropdown-menu-trigger-width)',
        }"
      >
        <UButton
          type="button"
          color="neutral"
          variant="ghost"
          :block="!collapsed"
          :square="collapsed"
          :aria-label="`打开${siteBrand}站点菜单`"
          :class="[
            'min-h-11 gap-2 px-1.5 data-[state=open]:bg-elevated',
            !collapsed && 'w-full justify-start',
            collapsed && 'aspect-square justify-center px-0',
          ]"
        >
          <span
            class="grid size-7 shrink-0 place-items-center rounded-md bg-primary/10 text-primary"
          >
            <UIcon name="i-tabler-book" class="size-4" />
          </span>
          <span
            v-if="!collapsed"
            class="min-w-0 truncate text-sm font-semibold text-highlighted"
          >
            {{ siteBrand }}
          </span>
          <UIcon
            v-if="!collapsed"
            name="i-tabler-chevrons-up-down"
            class="ms-auto size-3.5 text-dimmed"
          />
        </UButton>
      </UDropdownMenu>
    </template>

    <template #sidebar-footer="{ collapsed }">
      <ConsumerManageAccountControl
        home-to=""
        show-appearance
        :trigger-mode="collapsed ? 'collapsed' : 'sidebar'"
      />
    </template>

    <slot />
    <YBackToTop
      target-id="manage-main"
      scroll-container-id="manage-main"
      avoid-selector="[data-manage-dock], [data-back-to-top-avoid]"
      label="返回顶部"
    />
  </YAdminShell>
</template>
