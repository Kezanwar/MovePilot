// Form listing types
export interface FormListing {
  uuid: string;
  name: string;
  description: string | null;
  status: string;
  views: number;
  created_at: string; // ISO date string
  updated_at: string; // ISO date string
  affiliates: AffiliateInfo[];
  submission_count: number;
}

export interface AffiliateInfo {
  uuid: string;
  first_name: string;
  last_name: string;
}

// Full form types (for form builder/editor)
export interface Form {
  uuid: string;
  name: string;
  description: string | null;
  form_data: FormData;
  status: string;
  views: number;
  created_at: string;
  updated_at: string;
  affiliates: AffiliateInfo[];
  submission_count: number;
}

export interface FormData {
  steps: Step[];
  hero?: Hero;
}

export interface Step {
  uuid: string;
  title: string;
  description: string;
  inputs: Input[];
  condition?: Condition;
}

export interface Input {
  uuid: string;
  type: InputType;
  name: string;
  label: string;
  placeholder?: string;
  required: boolean;
  validation?: Validation;
  options?: Option[];
  default_value?: string;
  condition?: Condition;
}

export type InputType =
  | 'text'
  | 'email'
  | 'textarea'
  | 'select'
  | 'radio'
  | 'checkbox'
  | 'number'
  | 'date';

export interface Validation {
  // String validations
  min_length?: number;
  max_length?: number;
  length?: number;
  matches?: string;
  email?: boolean;
  url?: boolean;
  uuid?: boolean;

  // Number validations
  min?: number;
  max?: number;
  less_than?: number;
  more_than?: number;
  positive?: boolean;
  negative?: boolean;
  integer?: boolean;

  // Array validations
  min_items?: number;
  max_items?: number;

  // Date validations
  min_date?: string;
  max_date?: string;
}

export interface Option {
  uuid: string;
  label: string;
  value: string;
}

export interface Condition {
  operator: 'AND' | 'OR';
  conditions: CondRule[];
}

export interface CondRule {
  type: 'field' | 'step';
  uuid: string;
  operator: CondOperator;
  value: unknown;
}

export type CondOperator =
  | 'equals'
  | 'not_equals'
  | 'contains'
  | 'greater_than'
  | 'less_than'
  | 'is_empty'
  | 'is_not_empty'
  | 'is_completed';

export interface Hero {
  src: string;
  alt: string;
}

// Form submission types
export interface FormSubmission {
  uuid: string;
  full_name: string | null;
  email: string | null;
  submission_data: Record<string, unknown>;
  submitted_at: string;
}

// API Response types
export interface GetListingResponse {
  forms: FormListing[];
}

export interface GetFormResponse {
  form: Form;
}

export interface GetSubmissionsResponse {
  submissions: FormSubmission[];
}
