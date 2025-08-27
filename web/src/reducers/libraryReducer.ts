export interface LibraryState {
  novels: Array<{ id: string; title: string }>;
  filters: { query: string };
  page: number;
  pageSize: number;
}

export type LibraryAction =
  | { type: 'SET_NOVELS'; payload: Array<{ id: string; title: string }> }
  | { type: 'SET_FILTERS'; payload: { query: string } }
  | { type: 'SET_PAGE'; payload: number }
  | { type: 'SET_PAGE_SIZE'; payload: number };

export const initialLibraryState: LibraryState = {
  novels: [],
  filters: { query: '' },
  page: 1,
  pageSize: 10,
};

export function libraryReducer(state: LibraryState, action: LibraryAction): LibraryState {
  switch (action.type) {
    case 'SET_NOVELS':
      return { ...state, novels: action.payload };
    case 'SET_FILTERS':
      return { ...state, filters: action.payload, page: 1 };
    case 'SET_PAGE':
      return { ...state, page: action.payload };
    case 'SET_PAGE_SIZE':
      return { ...state, pageSize: action.payload, page: 1 };
    default:
      return state;
  }
}
