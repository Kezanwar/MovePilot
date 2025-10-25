import { makeObservable, observable, action, computed } from 'mobx';
import { RootStore } from '@app/stores/index';
import { clearSession, setSession } from '@app/lib/axios';
import { getInitialize } from '@app/api/auth';

import type { ManualAuthResponse } from '@app/api/auth';
import type { User } from '@app/types/user';

class AuthStore {
  rootStore: RootStore;

  constructor(rootStore: RootStore) {
    makeObservable(this, {
      user: observable,
      isAuthenticated: computed,
      isInitialized: observable,
      initialize: action,
      authenticate: action,
      unauthenticate: action
    });

    this.rootStore = rootStore;
  }

  user: User | undefined;

  get isAuthenticated() {
    return Boolean(this.user);
  }

  isInitialized = false;

  initialize = async () => {
    try {
      const res = await getInitialize();
      this.user = res.data.user;
    } catch (err) {
      console.error(err);
      this.unauthenticate();
    } finally {
      this.isInitialized = true;
    }
  };

  authenticate = (data: ManualAuthResponse) => {
    this.user = data.user;
    setSession(data.token);
  };

  unauthenticate = () => {
    clearSession();
    this.user = undefined;
  };
}

export default AuthStore;
