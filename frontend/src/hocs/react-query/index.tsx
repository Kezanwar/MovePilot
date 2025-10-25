import { type FC, type ReactNode } from 'react';
// import { ReactQueryDevtools } from '@tanstack/react-query-devtools';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
// import { IS_DEV } from '@app/config';

export const queryClient = new QueryClient();

type Props = { children: ReactNode };

const ReactQueryProvider: FC<Props> = ({ children }) => {
  return (
    <QueryClientProvider client={queryClient}>
      {/* {IS_DEV && <ReactQueryDevtools />} */}
      {children}
    </QueryClientProvider>
  );
};

export default ReactQueryProvider;
