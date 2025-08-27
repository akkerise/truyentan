import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { Link, useNavigate } from 'react-router-dom';
import { useContext } from 'react';
import { signInSchema, type SignInFormValues } from '../../forms/validators/auth';
import FormProvider from '../../forms/providers/FormProvider';
import { AuthContext } from '../../context/AuthContext';
import { post } from '../../services/api';

export default function SignIn() {
  const methods = useForm<SignInFormValues>({
    resolver: yupResolver(signInSchema),
    defaultValues: { email: '', password: '' },
  });

  const auth = useContext(AuthContext);
  const navigate = useNavigate();

  const onSubmit = async (data: SignInFormValues) => {
    try {
      const tokens = await post<{ access_token: string; refresh_token: string }>(
        '/auth/signin',
        data,
      );
      const authTokens = {
        accessToken: tokens.access_token,
        refreshToken: tokens.refresh_token,
      };
      localStorage.setItem('accessToken', authTokens.accessToken);
      localStorage.setItem('refreshToken', authTokens.refreshToken);
      auth?.signIn(null, authTokens);
      navigate('/');
    } catch (err) {
      console.error(err);
    }
  };

  const {
    register,
    formState: { errors },
  } = methods;

  return (
    <FormProvider methods={methods} onSubmit={onSubmit}>
      <div>
        <h1>Sign In</h1>
        <div>
          <label htmlFor="email">Email</label>
          <input id="email" type="email" {...register('email')} />
          {errors.email && <p>{errors.email.message}</p>}
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input id="password" type="password" {...register('password')} />
          {errors.password && <p>{errors.password.message}</p>}
        </div>
        <button type="submit">Sign In</button>
        <p>
          Don't have an account? <Link to="/signup">Sign Up</Link>
        </p>
      </div>
    </FormProvider>
  );
}
