import { useState } from 'react';

const useBoolean = (): [boolean, () => void, () => void] => {
  const [value, setValue] = useState<boolean>(false);

  const open = (): void => {
    setValue(true);
  };

  const close = (): void => {
    setValue(false);
  };

  return [value, open, close];
};

export default useBoolean;
