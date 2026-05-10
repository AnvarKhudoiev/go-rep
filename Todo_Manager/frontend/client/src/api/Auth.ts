const AUTH_URL = "http://localhost:8081"
 
export interface AuthResponse {
  token: string
}
 
export interface RegisterResponse {
  message: string
  user_id: number
}
 
export interface AuthError {
  error: string
}
 
export const register = async (username: string, password: string): Promise<RegisterResponse> => {
  const res = await fetch(`${AUTH_URL}/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  })
  const data = await res.json()
  if (!res.ok) throw data as AuthError
  return data
}
 
export const login = async (username: string, password: string): Promise<AuthResponse> => {
  const res = await fetch(`${AUTH_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  })
  const data = await res.json()
  if (!res.ok) throw data as AuthError
  return data
}

export const apiLogout = async (token: string) => {
  await fetch(`${AUTH_URL}/logout`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  })
}

export const getUserById = async (userId: number, token: string) => {
  const res = await fetch(`${AUTH_URL}/users/${userId}`, {
    headers: { Authorization: `Bearer ${token}` },
  })
  return res.json() // { id, username }
}
 
export const getUserIdFromToken = (token: string): number => {
  const payload = JSON.parse(atob(token.split(".")[1]))
  return payload.user_id
}
 