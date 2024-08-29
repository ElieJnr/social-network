import { create } from "zustand";

const useStore = create((set) => ({

  // State variables
  user: null,
  isAuthenticated: false,
  loading: false,

  // Actions
  setUser: (user) => set({ user, isAuthenticated: true }),
  logout: () => set({ user: null, isAuthenticated: false }),

  setLoading: (loading) => set({ loading }),
}));

export default useStore;
