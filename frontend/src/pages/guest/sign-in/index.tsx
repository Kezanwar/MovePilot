import { observer } from '@app/stores';
import LoginForm from './login-form';
import type { FC } from 'react';
import GuestLayout from '@app/layouts/guest';

const SignIn: FC = observer(() => {
  return (
    <GuestLayout>
      <LoginForm />
    </GuestLayout>
  );
});

export default SignIn;
