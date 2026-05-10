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

const data = {
  avatar: { avatar: "/avatars/shadcn.jpg" },
  teams: [
    { name: "Acme Inc", logo: <GalleryVerticalEndIcon />, plan: "Enterprise" },
    { name: "Acme Corp.", logo: <AudioLinesIcon />, plan: "Startup" },
    { name: "Evil Corp.", logo: <TerminalIcon />, plan: "Free" },
  ],
  projects: [
    { name: "Design Engineering", url: "#", icon: <FrameIcon /> },
    { name: "Sales & Marketing", url: "#", icon: <PieChartIcon /> },
    { name: "Travel", url: "#", icon: <MapIcon /> },
  ],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const navigate = useNavigate()
  const location = useLocation()

  const navMain = [
    {
      title: "Task Manager",
      url: "#",
      icon: <TerminalSquareIcon />,
      isActive: true,
      items: [
        {
          title: "Tasks",
          url: "/dashboard/tasks",
          isActive: location.pathname === "/dashboard/tasks",
          onClick: () => navigate("/dashboard/tasks"),
        },
        {
          title: "Add Task",
          url: "/dashboard/add-task",
          isActive: location.pathname === "/dashboard/add-task",
          onClick: () => navigate("/dashboard/add-task"),
        },
      ],
    },
    {
      title: "Profile",
      url: "#",
      icon: <UserIcon />,
      isActive: location.pathname === "/dashboard/profile",
      items: [
        {
          title: "Statistics",
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