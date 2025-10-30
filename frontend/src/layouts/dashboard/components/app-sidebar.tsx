import * as React from 'react';
import {
  Settings2,
  CreditCard,
  Layers,
  Users,
  UserPen
  // TriangleAlert
} from 'lucide-react';

import { NavMain } from '@app/layouts/dashboard/components/nav-main';
import NavAccount from '@app/layouts/dashboard/components/nav-account';
import NavUser from '@app/layouts/dashboard/components/nav-user';

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail
} from '@app/components/ui/sidebar';
import TopSection from './top-section';
import type { SidebarData } from '@app/types/sidebar';

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const data: SidebarData = React.useMemo(() => {
    return {
      navMain: [
        {
          title: 'Clients',
          url: 'forms/manage',
          icon: Layers,
          // iconClassName: 'text-indigo-400',
          isActive: true,
          items: [
            {
              title: 'Manage',
              url: 'clients/manage'
            },
            {
              title: 'Add New',
              url: 'clients/create'
            }
          ]
        },
        {
          title: 'Vendors',
          url: 'forms/manage',
          icon: Users,
          // iconClassName: 'text-orange-600',
          isActive: true,
          items: [
            {
              title: 'Manage',
              url: 'forms/manage'
            },
            {
              title: 'Add New',
              url: 'forms/create'
            }
          ]
        }
        // {
        //   title: 'Trades',
        //   url: '#',
        //   icon: PieChart,
        //   items: [
        //     {
        //       title: 'General',
        //       url: '#'
        //     },
        //     {
        //       title: 'Team',
        //       url: '#'
        //     },
        //     {
        //       title: 'Billing',
        //       url: '#'
        //     },
        //     {
        //       title: 'Limits',
        //       url: '#'
        //     }
        //   ]
        // },
        // {
        //   title: 'Strategies',
        //   url: '#',
        //   icon: SquareTerminal,
        //   items: [
        //     {
        //       title: 'Genesis',
        //       url: '#'
        //     },
        //     {
        //       title: 'Explorer',
        //       url: '#'
        //     },
        //     {
        //       title: 'Quantum',
        //       url: '#'
        //     }
        //   ]
        // },
        // {
        //   title: 'Reports',
        //   url: '#',
        //   icon: BookOpen,
        //   items: [
        //     {
        //       title: 'Introduction',
        //       url: '#'
        //     },
        //     {
        //       title: 'Get Started',
        //       url: '#'
        //     },
        //     {
        //       title: 'Tutorials',
        //       url: '#'
        //     },
        //     {
        //       title: 'Changelog',
        //       url: '#'
        //     }
        //   ]
        // }
      ],
      account: [
        {
          title: 'Billing',
          url: '#',
          icon: CreditCard
          // iconClassName: 'text-teal-600'
        },
        {
          title: 'Subscription',
          url: '#',
          icon: Settings2
          // iconClassName: 'text-pink-600'
        },
        {
          title: 'Profile',
          url: '#',
          icon: UserPen
          // iconClassName: 'text-blue-600'
        }
      ]
    };
  }, []);

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <TopSection />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        <NavAccount account={data.account} />
      </SidebarContent>
      <SidebarFooter className="mb-2">
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
