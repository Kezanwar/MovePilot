import { makeObservable, observable, action } from 'mobx';
import { RootStore } from '@app/stores/index';
import type { Form } from '@app/types/form';
import { getSingleForm } from '@app/api/form';
import { errorHandler, type ErrorObject } from '@app/lib/axios';

class FormStore {
  rootStore: RootStore;
  form: Form | null = null;
  isLoading = false;
  error: ErrorObject | null = null;
  isInitialized = false;

  constructor(rootStore: RootStore) {
    makeObservable(this, {
      form: observable,
      isLoading: observable,
      error: observable,
      reset: action,
      initializeEditForm: action,
      initializeNewForm: action,
      isInitialized: observable
    });

    this.rootStore = rootStore;
  }

  initializeEditForm = async (uuid: string) => {
    try {
      this.isLoading = true;
      const res = await getSingleForm(uuid);
      this.form = res.data.form;
    } catch (error) {
      errorHandler(error, (err) => {
        this.error = err;
      });
    } finally {
      this.isLoading = false;
      this.isInitialized = true;
    }
  };

  initializeNewForm = async (form: Form) => {
    this.form = form;
    this.isInitialized = true;
  };

  reset = () => {
    this.error = null;
    this.form = null;
    this.isLoading = false;
    this.isInitialized = false;
  };
}

export default FormStore;
