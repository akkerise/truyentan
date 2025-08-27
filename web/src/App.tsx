import AppRouter from './routes';
import { AuthProvider } from './context/AuthContext';
import { UIProvider } from './context/UIContext';
import { LibraryProvider } from './context/LibraryContext';

export default function App() {
  return (
    <AuthProvider>
      <UIProvider>
        <LibraryProvider>
          <AppRouter />
        </LibraryProvider>
      </UIProvider>
    </AuthProvider>
  );
}
