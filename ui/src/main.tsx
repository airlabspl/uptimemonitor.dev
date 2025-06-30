import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import './main.css'
import { AuthProvider } from './auth/auth-context'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { ProtectedRoutes } from './auth/protected-routes'
import { GuestRoutes } from './auth/guest-routes'
import HomePage from './pages/home'
import LoginPage from './pages/login'
import { Toaster } from './components/ui/sonner'
import RegisterPage from './pages/register'
import Dashboard from './pages/dashboard'
import { VerifiedRoutes } from './auth/verified-routes'
import VerifyPage from './pages/verify'
import { AppProvider } from './app/app-context'
import { SetupRoutes } from './app/setup-routes'
import SetupPage from './pages/setup'
import { ThemeProvider } from './theme/theme-context'
import ResetPasswordLink from './pages/reset-password-link'
import ResetPassword from './pages/reset-password'
import AppLayout from './layouts/app'
import Monitor from './pages/monitor'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider>
      <AppProvider>
        <AuthProvider>
          <BrowserRouter>
            <Routes>
              <Route element={<SetupRoutes />}>
                <Route path="/" element={<HomePage />} />
                <Route element={<ProtectedRoutes />}>
                  <Route path="/verify" element={<VerifyPage />} />
                  <Route element={<VerifiedRoutes />}>
                    <Route element={<AppLayout />}>
                      <Route path="/dashboard" element={<Dashboard />} />
                      <Route path="/m/:uuid" element={<Monitor />} />
                    </Route>
                  </Route>
                </Route>
                <Route element={<GuestRoutes />}>
                  <Route path="/login" element={<LoginPage />} />
                  <Route path="/register" element={<RegisterPage />} />
                  <Route path="/reset-password-link" element={<ResetPasswordLink />} />
                  <Route path="/reset-password/:token" element={<ResetPassword />} />
                </Route>
              </Route>
              <Route path="/setup" element={<SetupPage />} />
              <Route path="*" element={<div>404 Not Found</div>} />
            </Routes>
          </BrowserRouter>
        </AuthProvider>
      </AppProvider>
    </ThemeProvider>
    <Toaster position="bottom-left" />
  </StrictMode>,
)
