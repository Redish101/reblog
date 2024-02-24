import api from "@/api";

export function token() {
  return {
    headers: {
      Authorization: localStorage.getItem("token")!,
    },
  };
}

export async function checkAuth() {
  const data = await api.admin.tokenStateList(token());
  console.info("check auth", data.data.data);
  return data.data.data;
}
