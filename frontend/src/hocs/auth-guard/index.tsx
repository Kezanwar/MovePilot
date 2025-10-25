import store, { observer } from '@app/stores';
import { useEffect, type FC, type ReactNode } from 'react';
import { Navigate, useNavigate } from 'react-router';

type Props = {
  children: ReactNode;
};

const AuthGuard: FC<Props> = observer(({ children }) => {
  const nav = useNavigate();

  useEffect(() => {
    if (
      store.auth.isInitialized &&
      store.auth.isAuthenticated &&
      !store.auth.user?.email_confirmed
    ) {
      nav('/confirm-otp');
    }
  }, [
    nav,
    store.auth.isInitialized,
    store.auth.isAuthenticated,
    store.auth.user?.email_confirmed
  ]);

  if (store.auth.isInitialized) {
    if (!store.auth.isAuthenticated) {
      return (
        <Navigate to={'/sign-in'} state={{ to: window.location.pathname }} />
      );
    }
  }

  return children;
});

export default AuthGuard;
