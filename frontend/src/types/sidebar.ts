import type { LucideIcon } from 'lucide-react';

export interface NavItem {
  title: string;
  url?: string;
  method?: () => void;
  icon?: LucideIcon;
  iconClassName?: string;
}

export interface NavMain {
  title: string;
  url?: string;
  icon?: LucideIcon;
  iconClassName?: string;
  isActive?: boolean;
  items?: NavItem[];
}

export interface SidebarData {
  navMain: NavMain[];
  account: NavItem[];
}
