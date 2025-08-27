/* eslint-disable react-refresh/only-export-components */
import { createContext, useCallback, useMemo, useState, type ReactNode } from 'react';

interface AuthTokens {
  accessToken: string;
  refreshToken: string;
}

interface AuthContextType {
  user: unknown;
  tokens: AuthTokens | null;
  signIn: (user: unknown, tokens: AuthTokens) => void;
  signOut: () => void;
  refresh: (tokens: AuthTokens) => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<unknown>(null);
  const [tokens, setTokens] = useState<AuthTokens | null>(null);

  const signIn = useCallback((nextUser: unknown, nextTokens: AuthTokens) => {
    setUser(nextUser);
    setTokens(nextTokens);
  }, []);

  const signOut = useCallback(() => {
    setUser(null);
    setTokens(null);
  }, []);

  const refresh = useCallback((nextTokens: AuthTokens) => {
    setTokens(nextTokens);
  }, []);

  const value = useMemo(
    () => ({ user, tokens, signIn, signOut, refresh }),
    [user, tokens, signIn, signOut, refresh],
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
