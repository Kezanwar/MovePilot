import { observer } from '@app/stores';
import type { FC } from 'react';

import BasicLayout from '@app/layouts/basic';
import ConfirmOTPForm from './confirm-otp-form';

const ConfirmOTP: FC = observer(() => {
  return (
    <BasicLayout>
      <ConfirmOTPForm />
    </BasicLayout>
  );
});

export default ConfirmOTP;
