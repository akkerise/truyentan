import {
  FormProvider as RHFormProvider,
  UseFormReturn,
  FieldValues,
  SubmitHandler,
} from 'react-hook-form';
import { ReactNode } from 'react';

interface Props<T extends FieldValues> {
  methods: UseFormReturn<T>;
  onSubmit: SubmitHandler<T>;
  children: ReactNode;
}

export default function FormProvider<T extends FieldValues>({
  methods,
  onSubmit,
  children,
}: Props<T>) {
  return (
    <RHFormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>{children}</form>
    </RHFormProvider>
  );
}
