import { create } from 'zustand'
import { User } from '@/types/user'

interface UserStore {
  users: User[];
  isLoading: boolean;
  error: string | null;
  fetchUsers: (token: string) => Promise<void>;
}

export const useUserStore = create<UserStore>((set) => ({
  users: [],
  isLoading: false,
  error: null,

  fetchUsers: async (token: string) => {
    set({ isLoading: true, error: null });
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/user`, {
        headers: {
          'Accept': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });

      if (!response.ok) {
        throw new Error('Kullanıcılar yüklenirken bir hata oluştu');
      }

      const data = await response.json();
      set({ users: data, isLoading: false });
    } catch (error) {
      set({ error: error instanceof Error ? error.message : 'Kullanıcılar yüklenemedi', isLoading: false });
    }
  },
}));
