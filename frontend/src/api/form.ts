import axiosInstance from '@app/lib/axios';
import type { Form, FormListing } from '@app/types/form';

type FormListingResponse = {
  forms: FormListing[];
};

export const getFormListing = () =>
  axiosInstance.get<FormListingResponse>('/form/list');

type SingleFormResponse = {
  form: Form;
};

export const getSingleForm = (uuid: string) =>
  axiosInstance.get<SingleFormResponse>(`/form/view/${uuid}`);

export const createNewForm = () =>
  axiosInstance.post<SingleFormResponse>('/form/new');
