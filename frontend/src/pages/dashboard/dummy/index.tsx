import { Typography } from '@app/components/ui/typography';
import type { FC } from 'react';

type Props = {
  page: string;
};

const Dummy: FC<Props> = ({ page }) => {
  return <Typography>{page}</Typography>;
};

export default Dummy;
