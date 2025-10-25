import { Typography } from '@app/components/ui/typography';
import useFormListingQuery from '@app/hooks/queries/use-form-listing-query';
import { DataTable } from '@app/components/data-table';
import { SortableHeader, TableHeaderWithIcon } from '@app/components/ui/table';
import {
  ChartSpline,
  ToggleLeft,
  Users,
  Layers,
  Pencil,
  Trash,
  MoreVertical
} from 'lucide-react';
import { Button } from '@app/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger
} from '@app/components/ui/dropdown-menu';
import type { ColumnDef } from '@tanstack/react-table';
import {
  Avatar,
  AvatarFallback,
  getAvatarColorFromId,
  getInitials
} from '@app/components/ui/avatar';
import { observer } from 'mobx-react-lite';
import { useNavigate } from 'react-router';
import { useCallback } from 'react';
import SHSFUIButton from '@app/components/shsfui/button';

type Affiliate = {
  uuid: string;
  first_name: string;
  last_name: string;
};

type DummyData = {
  id: string;
  name: string;
  status: string;
  submissions: number;
  affiliates: Affiliate[];
};

const dummyData: DummyData[] = [
  {
    id: '1',
    name: 'Contact Form',
    status: 'Active',
    submissions: 42,
    affiliates: [
      { uuid: '1a', first_name: 'Kez', last_name: 'Anwar' },
      { uuid: '1b', first_name: 'Joe', last_name: 'Gill' }
    ]
  },
  {
    id: '2',
    name: 'Survey Form',
    status: 'Draft',
    submissions: 0,
    affiliates: [
      { uuid: '2a', first_name: 'Sarah', last_name: 'Smith' },
      { uuid: '2b1231312', first_name: 'Mike', last_name: 'Johnson' },
      { uuid: '2c', first_name: 'Emily', last_name: 'Davis' },
      { uuid: '231231', first_name: 'Emily', last_name: 'Davis' }
    ]
  },
  {
    id: '3',
    name: 'Feedback Form',
    status: 'Active',
    submissions: 128,
    affiliates: [{ uuid: '3a', first_name: 'Kez', last_name: 'Anwar' }]
  },
  {
    id: '4',
    name: 'Registration Form',
    status: 'Active',
    submissions: 87,
    affiliates: [
      { uuid: '4a', first_name: 'Joe', last_name: 'Gill' },
      { uuid: '4b', first_name: 'Alex', last_name: 'Brown' }
    ]
  },
  {
    id: '5',
    name: 'Application Form',
    status: 'Inactive',
    submissions: 15,
    affiliates: []
  }
];

const columns: ColumnDef<DummyData>[] = [
  {
    accessorKey: 'name',
    header: ({ column }) => (
      <SortableHeader
        column={column}
        icon={<Layers className="text-indigo-400" size={16} />}
      >
        Name
      </SortableHeader>
    )
  },
  {
    accessorKey: 'status',
    header: () => (
      <TableHeaderWithIcon
        icon={<ToggleLeft className="text-amber-700" size={16} />}
      >
        Status
      </TableHeaderWithIcon>
    )
  },
  {
    accessorKey: 'submissions',
    header: ({ column }) => (
      <SortableHeader
        column={column}
        icon={<ChartSpline className="text-teal-600" size={16} />}
      >
        Submissions
      </SortableHeader>
    )
  },
  {
    accessorKey: 'affiliates',
    header: () => (
      <TableHeaderWithIcon
        icon={<Users className="text-orange-600" size={16} />}
      >
        Affiliates
      </TableHeaderWithIcon>
    ),
    cell: ({ row }) => {
      const affiliates = row.original.affiliates;
      const maxDisplay = 3;
      const displayAffiliates = affiliates.slice(0, maxDisplay);
      const remaining = affiliates.length - maxDisplay;

      if (affiliates.length === 0) {
        return <span className="text-muted-foreground text-sm">None</span>;
      }

      return (
        <div className="flex items-center -space-x-2">
          {displayAffiliates.map((affiliate) => {
            const initials = getInitials(
              affiliate.first_name,
              affiliate.last_name
            );
            const colorClass = getAvatarColorFromId(affiliate.uuid);
            return (
              <Avatar
                key={affiliate.uuid}
                className="border-background size-8 border-2"
              >
                <AvatarFallback
                  className={`text-xs font-medium text-black ${colorClass}`}
                >
                  {initials}
                </AvatarFallback>
              </Avatar>
            );
          })}
          {remaining > 0 && (
            <Avatar className="border-background size-8 border-2">
              <AvatarFallback className="bg-muted text-foreground text-xs font-medium">
                +{remaining}
              </AvatarFallback>
            </Avatar>
          )}
        </div>
      );
    }
  },
  {
    id: 'actions',
    // header: () => <div className="text-right pr-2">Actions</div>,
    cell: () => {
      return (
        <div className="text-right">
          <DropdownMenu>
            <DropdownMenuTrigger>
              <Button variant="ghost" size="icon">
                <MoreVertical size={16} />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem onClick={() => console.log('Edit clicked')}>
                <Pencil size={16} />
                Edit
              </DropdownMenuItem>
              <DropdownMenuItem onClick={() => console.log('Delete clicked')}>
                <Trash size={16} />
                Delete
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      );
    },
    size: 80
  }
];

const ManageForms = observer(() => {
  const { data } = useFormListingQuery();

  console.log(data?.data);

  const nav = useNavigate();

  const onRowClick = useCallback((row: DummyData) => {
    nav(`forms/${row.id}/view`);
  }, []);

  return (
    <div className="space-y-4">
      <Typography variant={'h2'}>Manage Forms</Typography>
      {/* <SHSFUIButton>hello</SHSFUIButton> */}
      <DataTable onRowClick={onRowClick} columns={columns} data={dummyData} />
    </div>
  );
});

export default ManageForms;
