import { useTheme } from "next-themes"
import { SunIcon, MoonIcon } from "lucide-react"
import { SidebarMenu, SidebarMenuItem, SidebarMenuButton } from "@/shadcn/components/ui/sidebar"
 
export function ThemeToggle() {
  const { theme, setTheme } = useTheme()
 
  const isDark = theme === "dark"
 
  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <SidebarMenuButton
          onClick={() => setTheme(isDark ? "light" : "dark")}
          tooltip={isDark ? "Светлая тема" : "Тёмная тема"}
        >
          {isDark ? <SunIcon className="size-4" /> : <MoonIcon className="size-4" />}
          <span>{isDark ? "Light" : "Dark"}</span>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </SidebarMenu>
  )
}
 