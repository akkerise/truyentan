/* eslint-disable react-refresh/only-export-components */
import { createContext, useCallback, useMemo, useState, type ReactNode } from 'react';

interface UIContextType {
  modals: Record<string, boolean>;
  theme: 'light' | 'dark';
  fontSize: number;
  openModal: (name: string) => void;
  closeModal: (name: string) => void;
  toggleTheme: () => void;
  setFontSize: (size: number) => void;
}

export const UIContext = createContext<UIContextType | undefined>(undefined);

export function UIProvider({ children }: { children: ReactNode }) {
  const [modals, setModals] = useState<Record<string, boolean>>({});
  const [theme, setTheme] = useState<'light' | 'dark'>('light');
  const [fontSize, setFontSizeState] = useState<number>(16);

  const openModal = useCallback((name: string) => {
    setModals((m) => ({ ...m, [name]: true }));
  }, []);

  const closeModal = useCallback((name: string) => {
    setModals((m) => ({ ...m, [name]: false }));
  }, []);

  const toggleTheme = useCallback(() => {
    setTheme((t) => (t === 'light' ? 'dark' : 'light'));
  }, []);

  const setFontSize = useCallback((size: number) => {
    setFontSizeState(size);
  }, []);

  const value = useMemo(
    () => ({ modals, theme, fontSize, openModal, closeModal, toggleTheme, setFontSize }),
    [modals, theme, fontSize, openModal, closeModal, toggleTheme, setFontSize],
  );

  return <UIContext.Provider value={value}>{children}</UIContext.Provider>;
}
