import * as yup from 'yup';

export const signInSchema = yup.object({
  email: yup.string().email('Invalid email').required('Email is required'),
  password: yup.string().required('Password is required'),
});

export const signUpSchema = yup.object({
  name: yup.string().required('Name is required'),
  email: yup.string().email('Invalid email').required('Email is required'),
  password: yup.string().required('Password is required'),
});

export type SignInFormValues = yup.InferType<typeof signInSchema>;
export type SignUpFormValues = yup.InferType<typeof signUpSchema>;
