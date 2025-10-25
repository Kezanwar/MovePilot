import { observer } from '@app/stores';
import RegisterForm from './register-form';
import type { FC } from 'react';
import GuestLayout from '@app/layouts/guest';

const Register: FC = observer(() => {
  return (
    <GuestLayout>
      <RegisterForm />
    </GuestLayout>
  );
});

export default Register;
