import { create } from "zustand";

const useStore = create((set) => ({

  // State variables
  user: [],
  getMessage:[],
  isAuthenticated: false,
  loading: false,

  // Actions
  setUser: (user) => set({ user }),
  setgetMessage:(getMessage)=>set(getMessage),
  logout: () => set({ user: null, isAuthenticated: false }),

  setLoading: (loading) => set({ loading }),
}));

// const response = await fetch('https://api.example.com/data', {
//     next: { revalidate: 10 }, // Revalidate every 10 seconds
//   });

export default useStore;
