// 'use client'

// import { useState } from 'react'
// import { useRouter } from 'next/navigation'
// import { useForm } from 'react-hook-form'
// import { zodResolver } from '@hookform/resolvers/zod'
// import * as z from 'zod'
// import { Button } from '@/components/ui/button'
// import { Input } from '@/components/ui/input'
// import { Label } from '@/components/ui/label'
// import { Textarea } from '@/components/ui/textarea'
// import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
// import { Alert, AlertDescription } from '@/components/ui/alert'
// import { LockIcon, MailIcon, UserIcon, CalendarIcon, ImageIcon, AtSignIcon, FileTextIcon } from 'lucide-react'

// const loginSchema = z.object({
//   emailOrUsername: z.string().min(1, { message: 'Email or username is required' }),
//   password: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
// })

// const registerSchema = z.object({
//   email: z.string().email({ message: 'Invalid email address' }),
//   password: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
//   firstName: z.string().min(1, { message: 'First name is required' }),
//   lastName: z.string().min(1, { message: 'Last name is required' }),
//   dateOfBirth: z.string().refine((date) => !isNaN(Date.parse(date)), { message: 'Invalid date' }),
//   avatar: z.any().optional(),
//   nickname: z.string().optional(),
//   aboutMe: z.string().optional(),
// })

// export default function AuthForm() {
//   const router = useRouter()
//   const [isLogin, setIsLogin] = useState(true)
//   const [error, setError] = useState('')

//   const {
//     register,
//     handleSubmit,
//     formState: { errors, isSubmitting },
//     reset,
//   } = useForm({
//     resolver: zodResolver(isLogin ? loginSchema : registerSchema),
//   })

//   const onSubmit = async (data) => {
//     setError('')
//     try {
//       const response = await fetch('/api/auth', {
//         method: 'POST',
//         headers: {
//           'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({ ...data, action: isLogin ? 'login' : 'register' }),
//       })

//       if (!response.ok) {
//         throw new Error('Authentication failed')
//       }

//       const result = await response.json()
//       console.log(isLogin ? 'Login successful' : 'Registration successful', result)
//       router.push('/dashboard')
//     } catch (err) {
//       setError(isLogin ? 'Invalid email/username or password' : 'Registration failed')
//     }
//   }

//   const toggleForm = () => {
//     setIsLogin(!isLogin)
//     reset()
//     setError('')
//   }

//   return (
//     <Card className="w-full max-w-md">
//       <CardHeader>
//         <CardTitle className="text-2xl font-bold text-center">{isLogin ? 'Login' : 'Register'}</CardTitle>
//         <CardDescription className="text-center">
//           {isLogin ? 'Enter your credentials to access your account' : 'Create a new account'}
//         </CardDescription>
//       </CardHeader>
//       <form onSubmit={handleSubmit(onSubmit)}>
//         <CardContent className="space-y-4">
//           {isLogin ? (
//             <div className="space-y-2">
//               <Label htmlFor="emailOrUsername">Email or Username</Label>
//               <div className="relative">
//                 <AtSignIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                 <Input
//                   id="emailOrUsername"
//                   placeholder="Enter your email or username"
//                   className="pl-10"
//                   {...register('emailOrUsername')}
//                 />
//               </div>
//               {errors.emailOrUsername && <p className="text-sm text-red-500">{errors.emailOrUsername.message}</p>}
//             </div>
//           ) : (
//             <>
//               <div className="space-y-2">
//                 <Label htmlFor="email">Email</Label>
//                 <div className="relative">
//                   <MailIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                   <Input
//                     id="email"
//                     type="email"
//                     placeholder="Enter your email"
//                     className="pl-10"
//                     {...register('email')}
//                   />
//                 </div>
//                 {errors.email && <p className="text-sm text-red-500">{errors.email.message}</p>}
//               </div>
//               <div className="space-y-2">
//                 <Label htmlFor="firstName">First Name</Label>
//                 <div className="relative">
//                   <UserIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                   <Input
//                     id="firstName"
//                     placeholder="Enter your first name"
//                     className="pl-10"
//                     {...register('firstName')}
//                   />
//                 </div>
//                 {errors.firstName && <p className="text-sm text-red-500">{errors.firstName.message}</p>}
//               </div>
//               <div className="space-y-2">
//                 <Label htmlFor="lastName">Last Name</Label>
//                 <div className="relative">
//                   <UserIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                   <Input
//                     id="lastName"
//                     placeholder="Enter your last name"
//                     className="pl-10"
//                     {...register('lastName')}
//                   />
//                 </div>
//                 {errors.lastName && <p className="text-sm text-red-500">{errors.lastName.message}</p>}
//               </div>
//               <div className="space-y-2">
//                 <Label htmlFor="dateOfBirth">Date of Birth</Label>
//                 <div className="relative">
//                   <CalendarIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                   <Input
//                     id="dateOfBirth"
//                     type="date"
//                     className="pl-10"
//                     {...register('dateOfBirth')}
//                   />
//                 </div>
//                 {errors.dateOfBirth && <p className="text-sm text-red-500">{errors.dateOfBirth.message}</p>}
//               </div>
//               <div className="space-y-2">
//                 <Label htmlFor="avatar">Avatar/Image (Optional)</Label>
//                 <div className="relative">
//                   <ImageIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                   <Input
//                     id="avatar"
//                     type="file"
//                     className="pl-10"
//                     {...register('avatar')}
//                   />
//                 </div>
//               </div>
//               <div className="space-y-2">
//                 <Label htmlFor="nickname">Nickname (Optional)</Label>
//                 <div className="relative">
//                   <AtSignIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//                   <Input
//                     id="nickname"
//                     placeholder="Enter your nickname"
//                     className="pl-10"
//                     {...register('nickname')}
//                   />
//                 </div>
//               </div>
//               <div className="space-y-2">
//                 <Label htmlFor="aboutMe">About Me (Optional)</Label>
//                 <div className="relative">
//                   <FileTextIcon className="absolute left-3 top-3 text-gray-400" size={18} />
//                   <Textarea
//                     id="aboutMe"
//                     placeholder="Tell us about yourself"
//                     className="pl-10 min-h-[100px]"
//                     {...register('aboutMe')}
//                   />
//                 </div>
//               </div>
//             </>
//           )}
//           <div className="space-y-2">
//             <Label htmlFor="password">Password</Label>
//             <div className="relative">
//               <LockIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
//               <Input
//                 id="password"
//                 type="password"
//                 placeholder="Enter your password"
//                 className="pl-10"
//                 {...register('password')}
//               />
//             </div>
//             {errors.password && <p className="text-sm text-red-500">{errors.password.message}</p>}
//           </div>
//           {error && (
//             <Alert variant="destructive">
//               <AlertDescription>{error}</AlertDescription>
//             </Alert>
//           )}
//         </CardContent>
//         <CardFooter className="flex flex-col space-y-4">
//           <Button type="submit" className="w-full" disabled={isSubmitting}>
//             {isSubmitting ? (isLogin ? 'Logging in...' : 'Registering...') : (isLogin ? 'Log in' : 'Register')}
//           </Button>
//           <Button variant="link" type="button" onClick={toggleForm} className="w-full">
//             {isLogin ? 'Need an account? Register' : 'Already have an account? Log in'}
//           </Button>
//         </CardFooter>
//       </form>
//     </Card>
//   )
// }


'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import * as z from 'zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { LockIcon, MailIcon, UserIcon, CalendarIcon, ImageIcon, AtSignIcon, FileTextIcon } from 'lucide-react'

const loginSchema = z.object({
    emailOrUsername: z.string().min(1, { message: 'Email or username is required' }),
    password: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
})

const registerSchema = z.object({
    email: z.string().email({ message: 'Invalid email address' }),
    password: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
    firstname: z.string().min(1, { message: 'First name is required' }),
    lastname: z.string().min(1, { message: 'Last name is required' }),
    dateOfBirth: z.string().refine((date) => !isNaN(Date.parse(date)), { message: 'Invalid date' }),
    avatar: z.any().optional(),
    username: z.string().optional(),
    bio: z.string().optional(),
})

export default function AuthForm() {
    const router = useRouter()
    const [isLogin, setIsLogin] = useState(true)
    const [error, setError] = useState('')

    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
        reset,
    } = useForm({
        resolver: zodResolver(isLogin ? loginSchema : registerSchema),
    })

    const onSubmit = async (data) => {
        setError('');
        try {
            const url = isLogin ? "http://localhost:8080/login" : "http://localhost:8080/signin";
            
            const formData = new FormData();
            Object.keys(data).forEach((key) => {
                formData.append(key, data[key]);
            });
    
            // Ajouter le fichier avatar s'il est présent
            if (data.avatar && data.avatar.length > 0) {
                formData.append('avatar', data.avatar[0]);
            }

    
            const response = await fetch(url, {
                method: 'POST',
                body: formData,
            });
    
            if (!response.ok) {
                const errorMessage = await response.text();
                // throw new Error('Authentication failed');
                throw new Error(errorMessage);
            }
    
            const result = await response.json();
            console.log(isLogin ? 'Login successful' : 'Registration successful', result);
            router.push('/');
        } catch (err) {
            setError(isLogin ? 'Invalid email/username or password' : err.message);
        }
    };
    
    

    const toggleForm = () => {
        setIsLogin(!isLogin)
        reset()
        setError('')
    }

    return (
        <Card className="w-full max-w-md">
            <CardHeader>
                <CardTitle className="text-2xl font-bold text-center">{isLogin ? 'Login' : 'Register'}</CardTitle>
                <CardDescription className="text-center">
                    {isLogin ? 'Enter your credentials to access your account' : 'Create a new account'}
                </CardDescription>
            </CardHeader>
            <form onSubmit={handleSubmit(onSubmit)}>
                <CardContent className="space-y-4">
                    {isLogin ? (
                        <div className="space-y-2">
                            <Label htmlFor="emailOrUsername">Email or Username</Label>
                            <div className="relative">
                                <AtSignIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                <Input
                                    id="emailOrUsername"
                                    placeholder="Enter your email or username"
                                    className="pl-10"
                                    {...register('emailOrUsername')}
                                />
                            </div>
                            {errors.emailOrUsername && <p className="text-sm text-red-500">{errors.emailOrUsername.message}</p>}
                        </div>
                    ) : (
                        <>
                            <div className="grid grid-cols-2 gap-4">
                                <div className="space-y-2">
                                    <Label htmlFor="firstname">First Name</Label>
                                    <div className="relative">
                                        <UserIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                        <Input
                                            id="firstname"
                                            placeholder="Enter your first name"
                                            className="pl-10"
                                            {...register('firstname')}
                                        />
                                    </div>
                                    {errors.firstname && <p className="text-sm text-red-500">{errors.firstname.message}</p>}
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="lastname">Last Name</Label>
                                    <div className="relative">
                                        <UserIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                        <Input
                                            id="lastname"
                                            placeholder="Enter your last name"
                                            className="pl-10"
                                            {...register('lastname')}
                                        />
                                    </div>
                                    {errors.lastname && <p className="text-sm text-red-500">{errors.lastname.message}</p>}
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="email">Email</Label>
                                    <div className="relative">
                                        <MailIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                        <Input
                                            id="email"
                                            type="email"
                                            placeholder="Enter your email"
                                            className="pl-10"
                                            {...register('email')}
                                        />
                                    </div>
                                    {errors.email && <p className="text-sm text-red-500">{errors.email.message}</p>}
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="username">Username (Optional)</Label>
                                    <div className="relative">
                                        <AtSignIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                        <Input
                                            id="username"
                                            placeholder="Enter your uername"
                                            className="pl-10"
                                            {...register('username')}
                                        />
                                    </div>
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="avatar">Avatar/Image (Optional)</Label>
                                    <div className="relative">
                                        <ImageIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                        <Input
                                            id="avatar"
                                            type="file"
                                            className="pl-10"
                                            {...register('avatar')}
                                        />
                                    </div>
                                </div>
                                <div className="space-y-2">
                                    <Label htmlFor="dateOfBirth">Date of Birth</Label>
                                    <div className="relative">
                                        <CalendarIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                                        <Input
                                            id="dateOfBirth"
                                            type="date"
                                            className="pl-10"
                                            {...register('dateOfBirth')}
                                        />
                                    </div>
                                    {errors.dateOfBirth && <p className="text-sm text-red-500">{errors.dateOfBirth.message}</p>}
                                </div>
                            </div>

                            <div className="space-y-2">
                                <Label htmlFor="bio">About Me (Optional)</Label>
                                <div className="relative">
                                    <FileTextIcon className="absolute left-3 top-3 text-gray-400" size={18} />
                                    <Textarea
                                        id="bio"
                                        placeholder="Tell us about yourself"
                                        className="pl-10 min-h-[100px]"
                                        {...register('bio')}
                                    />
                                </div>
                            </div>
                        </>
                    )}
                    <div className="space-y-2">
                        <Label htmlFor="password">Password</Label>
                        <div className="relative">
                            <LockIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={18} />
                            <Input
                                id="password"
                                type="password"
                                placeholder="Enter your password"
                                className="pl-10"
                                {...register('password')}
                            />
                        </div>
                        {errors.password && <p className="text-sm text-red-500">{errors.password.message}</p>}
                    </div>
                    {error && (
                        <Alert variant="destructive">
                            <AlertDescription>{error}</AlertDescription>
                        </Alert>
                    )}
                </CardContent>
                <CardFooter className="flex flex-col space-y-4">
                    <Button type="submit" className="w-full" disabled={isSubmitting}>
                        {isSubmitting ? (isLogin ? 'Logging in...' : 'Registering...') : (isLogin ? 'Log in' : 'Register')}
                    </Button>
                    <Button variant="link" type="button" onClick={toggleForm} className="w-full">
                        {isLogin ? 'Need an account? Register' : 'Already have an account? Log in'}
                    </Button>
                </CardFooter>
            </form>
        </Card>
    )
}
