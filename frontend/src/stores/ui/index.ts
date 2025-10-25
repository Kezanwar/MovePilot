import { observable, action, makeObservable } from 'mobx';
import type { RootStore } from '..';

type Theme = 'light' | 'dark';

class UIStore {
  rootStore: RootStore;
  theme: Theme = 'dark';
  isLoading: number = 0;

  constructor(rootStore: RootStore) {
    makeObservable(this, {
      theme: observable,
      setTheme: action,
      isLoading: observable,
      addLoading: action,
      removeLoading: action
    });

    this.rootStore = rootStore;

    const saved = localStorage.getItem('$MobX-theme') as Theme;
    if (saved) {
      this.setTheme(saved);
    }

    document.documentElement.classList.remove('light', 'dark');
    document.documentElement.classList.add(this.theme);
  }

  setTheme(theme: Theme) {
    this.theme = theme;
    document.documentElement.classList.remove('light', 'dark');
    document.documentElement.classList.add(theme);
    localStorage.setItem('$MobX-theme', theme);
  }

  addLoading() {
    this.isLoading = this.isLoading + 1;
  }

  removeLoading() {
    this.isLoading = Math.max(0, this.isLoading - 1);
  }
}

export default UIStore;
