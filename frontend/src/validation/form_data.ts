// import * as Yup from 'yup';

// function buildYupSchema(inputs) {
//   const shape = {};

//   inputs.forEach((input) => {
//     let schema;

//     // Start with base type
//     switch (input.type) {
//       case 'email':
//       case 'text':
//       case 'textarea':
//         schema = Yup.string();
//         break;
//       case 'number':
//         schema = Yup.number();
//         break;
//       case 'date':
//         schema = Yup.date();
//         break;
//       case 'checkbox':
//         schema = Yup.array();
//         break;
//       case 'select':
//       case 'radio':
//         schema = Yup.string();
//         break;
//       default:
//         schema = Yup.string();
//     }

//     // Apply required
//     if (input.required) {
//       schema = schema.required();
//     }

//     // Apply validations if they exist
//     if (input.validation) {
//       const v = input.validation;

//       // String validations
//       if (v.minLength) schema = schema.min(v.minLength);
//       if (v.maxLength) schema = schema.max(v.maxLength);
//       if (v.length) schema = schema.length(v.length);
//       if (v.matches) schema = schema.matches(new RegExp(v.matches));
//       if (v.email) schema = schema.email();
//       if (v.url) schema = schema.url();
//       if (v.uuid) schema = schema.uuid();

//       // Number validations
//       if (v.min !== undefined) schema = schema.min(v.min);
//       if (v.max !== undefined) schema = schema.max(v.max);
//       if (v.lessThan !== undefined) schema = schema.lessThan(v.lessThan);
//       if (v.moreThan !== undefined) schema = schema.moreThan(v.moreThan);
//       if (v.positive) schema = schema.positive();
//       if (v.negative) schema = schema.negative();
//       if (v.integer) schema = schema.integer();

//       // Array validations
//       if (v.minItems) schema = schema.min(v.minItems);
//       if (v.maxItems) schema = schema.max(v.maxItems);

//       // Date validations
//       if (v.minDate) schema = schema.min(new Date(v.minDate));
//       if (v.maxDate) schema = schema.max(new Date(v.maxDate));
//     }

//     shape[input.name] = schema;
//   });

//   return Yup.object().shape(shape);
// }

// // Usage
// const formData = await fetchForm(formId);
// const validationSchema = buildYupSchema(
//   formData.steps.flatMap((s) => s.inputs)
// );

// function evaluateOperator(
//   fieldValue: any,
//   operator: CondOperator,
//   targetValue: any
// ): boolean {
//   switch (operator) {
//     case 'equals':
//       return fieldValue === targetValue;
//     case 'not_equals':
//       return fieldValue !== targetValue;
//     case 'contains':
//       return fieldValue?.includes(targetValue);
//     case 'greater_than':
//       return fieldValue > targetValue;
//     case 'less_than':
//       return fieldValue < targetValue;
//     case 'is_empty':
//       return !fieldValue || fieldValue === '';
//     case 'is_not_empty':
//       return !!fieldValue && fieldValue !== '';
//     case 'is_completed':
//       return fieldValue === true;
//     default:
//       return true;
//   }
// }
