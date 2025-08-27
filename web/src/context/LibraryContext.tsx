/* eslint-disable react-refresh/only-export-components */
import { createContext, useMemo, useReducer, type ReactNode } from 'react';
import {
  libraryReducer,
  initialLibraryState,
  type LibraryAction,
  type LibraryState,
} from '../reducers/libraryReducer';

interface LibraryContextType extends LibraryState {
  dispatch: React.Dispatch<LibraryAction>;
  filteredNovels: Array<{ id: string; title: string }>;
}

export const LibraryContext = createContext<LibraryContextType | undefined>(undefined);

export function LibraryProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(libraryReducer, initialLibraryState);

  const filteredNovels = useMemo(() => {
    const query = state.filters.query.toLowerCase();
    return state.novels.filter((n) => n.title.toLowerCase().includes(query));
  }, [state.novels, state.filters]);

  const value = useMemo(
    () => ({ ...state, dispatch, filteredNovels }),
    [state, dispatch, filteredNovels],
  );

  return <LibraryContext.Provider value={value}>{children}</LibraryContext.Provider>;
}
