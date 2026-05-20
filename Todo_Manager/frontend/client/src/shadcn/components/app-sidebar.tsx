import * as React from "react"
import { useNavigate, useLocation } from "react-router-dom"
import { NavMain } from "@/shadcn/components/nav-main"
import { NavProjects } from "@/shadcn/components/nav-projects"
import { NavUser } from "@/shadcn/components/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/shadcn/components/ui/sidebar"
import {
  GalleryVerticalEndIcon,
  AudioLinesIcon,
  TerminalIcon,
  TerminalSquareIcon,
  UserIcon,
  FrameIcon,
  PieChartIcon,
  MapIcon,
} from "lucide-react"
import { ThemeToggle } from "@/components/ThemeToggle"
import { useTodos } from "@/store/todoStore"

const data = {
  avatar: { avatar: "/avatars/shadcn.jpg" },
  teams: [
    { name: "Acme Inc", logo: <GalleryVerticalEndIcon />, plan: "Enterprise" },
    { name: "Acme Corp.", logo: <AudioLinesIcon />, plan: "Startup" },
    { name: "Evil Corp.", logo: <TerminalIcon />, plan: "Free" },
  ],
  projects: [
    { name: "Дизайн и разработка", url: "#", icon: <FrameIcon /> },
    { name: "Продажи и маркетинг", url: "#", icon: <PieChartIcon /> },
    { name: "Путешествия", url: "#", icon: <MapIcon /> },
  ],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const navigate = useNavigate()
  const location = useLocation()
  const { todos } = useTodos()

  const now = new Date()
  const overdueCount = todos.filter(t =>
    !t.completed && t.due_date && new Date(t.due_date) < now
  ).length

  const navMain = [
    {
      title: "Менеджер задач",
      url: "#",
      icon: <TerminalSquareIcon />,
      isActive: true,
      items: [
        {
          title: "Задачи",
          url: "/dashboard/tasks",
          isActive: location.pathname === "/dashboard/tasks",
          onClick: () => navigate("/dashboard/tasks"),
        },
        {
          title: "Новая задача",
          url: "/dashboard/add-task",
          isActive: location.pathname === "/dashboard/add-task",
          onClick: () => navigate("/dashboard/add-task"),
        },
        ...(overdueCount > 0 ? [{
          title: "Просроченные",
          url: "/dashboard/overdue",
          isActive: location.pathname === "/dashboard/overdue",
          onClick: () => navigate("/dashboard/overdue"),
          badge: overdueCount,
        }] : []),
      ],
    },
    {
      title: "Профиль",
      url: "#",
      icon: <UserIcon />,
      isActive: location.pathname === "/dashboard/profile",
      items: [
        {
          title: "Статистика",
          url: "/dashboard/profile",
          isActive: location.pathname === "/dashboard/profile",
          onClick: () => navigate("/dashboard/profile"),
        },
      ],
    },
  ]

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <NavUser avatar={data.avatar} />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={navMain} />
        <NavProjects projects={data.projects} />
      </SidebarContent>
      <SidebarFooter>
        <ThemeToggle />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}