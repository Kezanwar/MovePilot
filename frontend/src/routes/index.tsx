import React from 'react';
import { useRoutes, type RouteObject } from 'react-router-dom';

import Signin from '@app/pages/guest/sign-in';
import GuestGuard from '@app/hocs/guest-guard';
import DashboardLayout from '@app/layouts/dashboard';
import Home from '@app/pages/dashboard/home';
import Dummy from '@app/pages/dashboard/dummy';

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
        path: 'forms/new',
        element: (
          <Dummy page="Generate new form then redirect to /forms/<uuid>/create" />
        )
      }
    ]
  },
  {
    path: '/sign-in',
    element: (
      <GuestGuard>
        <Signin />
      </GuestGuard>
    )
  }
];

const Routes: React.FC = () => {
  const elements = useRoutes(paths);
  return elements;
};

export default Routes;
