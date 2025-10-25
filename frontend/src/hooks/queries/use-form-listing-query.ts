import { getFormListing } from '@app/api/form';
import { useQuery } from '@tanstack/react-query';

const useFormListingQuery = () => {
  return useQuery({
    queryKey: ['FORM_LISTING'],
    queryFn: getFormListing,
    staleTime: 15 * (60 * 1000) // 15mins
  });
};

export default useFormListingQuery;
