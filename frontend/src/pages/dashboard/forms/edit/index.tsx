import { Typography } from '@app/components/ui/typography';
import store from '@app/stores';
import { observer } from 'mobx-react-lite';
import { useEffect } from 'react';
import { useParams } from 'react-router';

const EditForm = observer(() => {
  const { uuid } = useParams();

  useEffect(() => {
    if (uuid && !store.form.isInitialized && !store.form.isLoading) {
      store.form.initializeEditForm(uuid);
    }
  }, [uuid, store.form.isInitialized, store.form.isLoading]);

  return (
    <div className="space-y-4">
      <Typography variant={'h2'}>Edit Form</Typography>
    </div>
  );
});

export default EditForm;
