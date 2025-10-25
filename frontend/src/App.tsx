import AuthInitializer from './hocs/auth-initializer';
import Toast from './components/toast';
import Routes from './routes';
import LoadingOverlay from './components/loading-overlay';
import ReactQueryProvider from './hocs/react-query';

function App() {
  return (
    <ReactQueryProvider>
      <AuthInitializer>
        <Toast />
        <LoadingOverlay />
        <Routes />
      </AuthInitializer>
    </ReactQueryProvider>
  );
}

export default App;
