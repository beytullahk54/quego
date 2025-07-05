import { create } from 'zustand'
import { Job } from '@/types/job'

interface UserStore {
  jobs: Job[];
  isLoading: boolean;
  error: string | null;
  fetchUsers: (token: string) => Promise<void>;
}

export const useUserStore = create<UserStore>((set) => ({
  jobs: [],
  isLoading: false,
  error: null,

  fetchUsers: async (token: string) => {
    set({ isLoading: true, error: null });
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/jobs`, {
        method: 'GET',
        headers: {
          'Accept': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });

      if (!response.ok) {
        throw new Error('İşler yüklenirken bir hata oluştu');
      }

      const data = await response.json();
      set({ jobs: data, isLoading: false });
    } catch (error) {
      set({ error: error instanceof Error ? error.message : 'İşler yüklenemedi', isLoading: false });
    }
  },
}));
