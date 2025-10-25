import * as React from 'react';
import {
  Settings2,
  CreditCard,
  Layers,
  Users,
  UserPen,
  TriangleAlert
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
import { createNewForm } from '@app/api/form';
import { errorHandler } from '@app/lib/axios';
import { toast } from 'sonner';
import store from '@app/stores';
import { useNavigate, type NavigateFunction } from 'react-router';

const handleCreateNewForm = async (nav: NavigateFunction) => {
  try {
    store.ui.addLoading();
    const res = await createNewForm();
    store.form.initializeNewForm(res.data.form);
    nav(`forms/${res.data.form.uuid}/edit`);
  } catch (error) {
    errorHandler(error, (e) =>
      toast(e.message, {
        position: 'bottom-left',
        icon: <TriangleAlert className="text-destructive mr-10" />,
        description:
          'Please try again, or contact support if further help is needed.'
      })
    );
  } finally {
    store.ui.removeLoading();
  }
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const nav = useNavigate();
  const data: SidebarData = React.useMemo(() => {
    return {
      navMain: [
        {
          title: 'Forms',
          url: 'forms/manage',
          icon: Layers,
          iconClassName: 'text-indigo-400',
          isActive: true,
          items: [
            {
              title: 'Manage',
              url: 'forms/manage'
            },
            {
              title: 'Add New',
              method: () => handleCreateNewForm(nav)
            }
          ]
        },
        {
          title: 'Affiliates',
          url: 'forms/manage',
          icon: Users,
          iconClassName: 'text-orange-600',
          isActive: true,
          items: [
            {
              title: 'Manage',
              url: 'forms/manage'
            },
            {
              title: 'Add',
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
          icon: CreditCard,
          iconClassName: 'text-teal-600'
        },
        {
          title: 'Subscription',
          url: '#',
          icon: Settings2,
          iconClassName: 'text-pink-600'
        },
        {
          title: 'Profile',
          url: '#',
          icon: UserPen,
          iconClassName: 'text-blue-600'
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
