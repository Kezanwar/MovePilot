import { cn } from '@app/lib/utils';
import { Button } from '@app/components/ui/button';
import { Card, CardContent, CardTitle } from '@app/components/ui/card';
import { ChevronRight, TriangleAlert } from 'lucide-react';
import { useForm, FormProvider } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { LoginSchema, type TLoginForm } from '@app/validation/auth';

import RHFInput from '@app/components/hook-form/rhf-input';
import type { FC } from 'react';
import { postSignIn } from '@app/api/auth';
import store from '@app/stores';
import { useLocation, useNavigate } from 'react-router';

import { errorHandler } from '@app/lib/axios';

import { toast } from 'sonner';
import { Link } from 'react-router-dom';
import SHSFUIButton from '@app/components/shsfui/button';

const LoginForm: FC = ({ ...props }) => {
  const nav = useNavigate();

  const { state } = useLocation();

  const methods = useForm<TLoginForm>({
    resolver: yupResolver(LoginSchema),
    defaultValues: {
      email: '',
      password: ''
    },
    mode: 'onSubmit'
  });

  const onSubmit = async (data: TLoginForm) => {
    try {
      store.ui.addLoading();
      const res = await postSignIn(data);
      store.auth.authenticate(res.data);
      nav(state?.to || '/');
    } catch (error) {
      errorHandler(error, (e) =>
        toast(e.message, {
          position: 'bottom-left',
          icon: <TriangleAlert className="text-destructive mr-10" />,
          description:
            'Please try again, or contact support if further help is needed.'
        })
      );
      store.auth.unauthenticate();
    } finally {
      store.ui.removeLoading();
    }
  };

  return (
    <FormProvider {...methods}>
      <Card>
        <CardContent>
          <CardTitle className="mb-2">Welcome Back</CardTitle>
          <div className="text-muted-foreground *:[a]:hover:text-primary text-left text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
            New here? <Link to={'/register'}>Sign up</Link>
          </div>
          <div className={cn('mt-8 flex flex-col gap-6')} {...props}>
            <form onSubmit={methods.handleSubmit(onSubmit)}>
              <div className="flex flex-col gap-5">
                <div className="grid gap-3">
                  <RHFInput
                    name="email"
                    label="Email"
                    placeholder="johndoe@example.com"
                  />
                  <RHFInput name="password" label="Password" type="password" />
                </div>
                <SHSFUIButton
                  icon={<ChevronRight />}
                  type="submit"
                  className="w-full"
                >
                  Login
                </SHSFUIButton>
              </div>

              <div className="after:border-border relative my-6 text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t">
                <span className="text-muted-foreground relative z-10 px-2">
                  Or
                </span>
              </div>

              <div className="grid gap-4 sm:grid-cols-2">
                <Button variant="outline" type="button" className="w-full">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path
                      d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.676-1.48 3.676-2.948 1.156-1.688 1.636-3.325 1.662-3.415-.039-.013-3.182-1.221-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.61 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.559-1.701"
                      fill="currentColor"
                    />
                  </svg>
                  Continue with Apple
                </Button>
                <Button variant="outline" type="button" className="w-full">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path
                      d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
                      fill="currentColor"
                    />
                  </svg>
                  Continue with Google
                </Button>
              </div>
            </form>

            <div className="text-muted-foreground *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
              By signing in, you agree to our <a href="#">Terms of Service</a>{' '}
              and <a href="#">Privacy Policy</a>.
            </div>
          </div>
        </CardContent>
      </Card>
    </FormProvider>
  );
};

export default LoginForm;
