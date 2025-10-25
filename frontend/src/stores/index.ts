import { observer } from 'mobx-react-lite';
import AuthStore from './auth';
import UIStore from './ui';
import FormStore from './form';

export class RootStore {
  auth = new AuthStore(this);
  ui = new UIStore(this);
  form = new FormStore(this);
}

const store = new RootStore();

export default store;

export { observer };
