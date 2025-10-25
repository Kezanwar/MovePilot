import React from 'react';
import { useRoutes, type RouteObject } from 'react-router-dom';

import Register from '@app/pages/guest/register';
import Signin from '@app/pages/guest/sign-in';
import GuestGuard from '@app/hocs/guest-guard';
import DashboardLayout from '@app/layouts/dashboard';
import Home from '@app/pages/dashboard/home';
import Dummy from '@app/pages/dashboard/dummy';
import AuthGuard from '@app/hocs/auth-guard';

// import BasicLayout from '@app/layouts/basic';
// import GuestLayout from '@app/layouts/guest';
import ConfirmOTP from '@app/pages/basic/confirm-otp';
import ManageForms from '@app/pages/dashboard/forms/manage';
import EditForm from '@app/pages/dashboard/forms/edit';

const paths: RouteObject[] = [
  {
    path: '/',
    element: <DashboardLayout />,
    children: [
      {
        index: true,
        element: <Home />
      },
      {
        path: 'forms/:uuid/edit',
        element: <EditForm />
      },
      {
        path: 'forms/new',
        element: (
          <Dummy page="Generate new form then redirect to /forms/<uuid>/create" />
        )
      },
      {
        path: 'forms/manage',
        element: <ManageForms />
      }
    ]
  },
  {
    path: '/confirm-otp',
    element: (
      <AuthGuard>
        <ConfirmOTP />
      </AuthGuard>
    )
  },
  {
    path: '/sign-in',
    element: (
      <GuestGuard>
        <Signin />
      </GuestGuard>
    )
  },
  {
    path: '/register',
    element: (
      <GuestGuard>
        <Register />
      </GuestGuard>
    )
  }
];

const Routes: React.FC = () => {
  const elements = useRoutes(paths);
  return elements;
};

export default Routes;
