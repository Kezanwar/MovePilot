import store, { observer } from '@app/stores';

import { toast } from 'sonner';

import { useNavigate } from 'react-router';
import { errorHandler } from '@app/lib/axios';
import {
  InputOTP,
  InputOTPGroup,
  InputOTPSeparator,
  InputOTPSlot
} from '@app/components/ui/input-otp';
import { Typography } from '@app/components/ui/typography';
import { Button } from '@app/components/ui/button';
import { useEffect, useState } from 'react';
import { confirmOTP } from '@app/api/auth';
import { TriangleAlert } from 'lucide-react';
import { Card, CardContent } from '@app/components/ui/card';

const ConfirmOTPForm = observer(() => {
  const nav = useNavigate();

  const [otp, setOtp] = useState('');

  const onInputChange = async (newValue: string) => {
    setOtp(newValue);
  };

  const handleSubmitOTP = async () => {
    try {
      store.ui.addLoading();
      const res = await confirmOTP(otp);
      store.auth.authenticate(res.data);
      nav('/');
    } catch (error) {
      errorHandler(error, (e) =>
        toast(e.message, {
          position: 'bottom-left',
          icon: <TriangleAlert className="text-destructive mr-10" />,
          description:
            'Please try again, or contact support if further help is needed.'
        })
      );
    } finally {
      store.ui.removeLoading();
    }
  };

  const onFormSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();
    handleSubmitOTP();
  };

  useEffect(() => {
    if (otp.length === 6) {
      handleSubmitOTP();
    }
  }, [otp]);

  return (
    <div className="mt-30 flex flex-col items-center justify-center gap-4 text-center">
      <Card className="py-8">
        <CardContent>
          <Typography className="mb-3" variant={'h2'}>
            Hello {store.auth.user?.first_name}
          </Typography>
          <div className="mb-8 flex w-120 flex-col">
            <Typography color="muted" className="text-center" variant={'small'}>
              Thanks for registering, before you can continue onto the Platform
              you must confirm your email address. An email with a 6 digit OTP
              code was sent to {store.auth.user?.email || ''}
            </Typography>
            <Typography color="muted" variant={'small'}></Typography>
          </div>
          <form
            onSubmit={onFormSubmit}
            className="mt-4 flex flex-col items-center justify-center gap-5"
          >
            <InputOTP onChange={onInputChange} maxLength={6}>
              <InputOTPGroup>
                <InputOTPSlot index={0} />
                <InputOTPSlot index={1} />
                <InputOTPSlot index={2} />
              </InputOTPGroup>
              <InputOTPSeparator />
              <InputOTPGroup>
                <InputOTPSlot index={3} />
                <InputOTPSlot index={4} />
                <InputOTPSlot index={5} />
              </InputOTPGroup>
            </InputOTP>
          </form>

          <div className="text-muted-foreground text-center text-xs text-balance">
            Didn't receive an email?{' '}
            <Button
              className="text-muted-foreground p-0 text-xs"
              variant={'link'}
            >
              Resend OTP Email
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
});

export default ConfirmOTPForm;
