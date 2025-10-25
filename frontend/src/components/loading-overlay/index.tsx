import { observer } from 'mobx-react-lite';
import { Spinner } from '@app/components/ui/spinner';
import store from '@app/stores';
import { cn } from '@app/lib/utils';
import { Typography } from '../ui/typography';

const LoadingOverlay = observer(() => {
  return store.ui.isLoading > 0 ? (
    <div className={cn('fixed inset-0 z-50')}>
      <div className="bg-background/60 absolute inset-0 opacity-70 backdrop-blur-xs" />
      <div className="absolute right-4 bottom-4 flex items-center gap-2">
        <Typography className="text-sm" variant={'small'}>
          Loading...
        </Typography>
        <Spinner size={32} className="text-foreground" />
      </div>
    </div>
  ) : null;
});

export default LoadingOverlay;
