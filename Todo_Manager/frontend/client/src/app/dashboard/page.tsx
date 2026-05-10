import { Routes, Route, Navigate, useNavigate, useLocation } from "react-router-dom"
import { AppSidebar } from "@/shadcn/components/app-sidebar"
import { Separator } from "@/shadcn/components/ui/separator"
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/shadcn/components/ui/sidebar"
import { TodoProvider } from "@/store/todoStore"
import { TodoList } from "@/components/todos/TodoList"
import { AddTodoPage } from "@/components/todos/AddTodoPage"
import { ProfilePage } from "@/components/profile/ProfilePage"

export default function Page() {
  const navigate = useNavigate()
  const location = useLocation()

  const getTitle = () => {
    if (location.pathname.includes("add-task")) return "Новая задача"
    if (location.pathname.includes("profile")) return "Профиль"
    return "Задачи"
  }

  return (
    <TodoProvider>
      <SidebarProvider>
        <AppSidebar />
        <SidebarInset>
          <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12">
            <div className="flex items-center gap-2 px-4">
              <SidebarTrigger className="-ml-1" />
              <Separator orientation="vertical" className="mr-2 data-[orientation=vertical]:h-4" />
              <h1 className="text-sm font-medium">{getTitle()}</h1>
            </div>
          </header>
          <div className="flex flex-1 flex-col p-6 pt-2">
            <Routes>
              <Route index element={<Navigate to="tasks" replace />} />
              <Route
                path="tasks"
                element={<TodoList onAddClick={() => navigate("/dashboard/add-task")} />}
              />
              <Route
                path="add-task"
                element={<AddTodoPage onSuccess={() => navigate("/dashboard/tasks")} />}
              />
              <Route path="profile" element={<ProfilePage />} />
            </Routes>
          </div>
        </SidebarInset>
      </SidebarProvider>
    </TodoProvider>
  )
}