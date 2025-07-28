import { Suspense } from "react";
import { BrowserRouter, Route, Routes } from "react-router";
import LoadingPage from "../../components/LoadingPage";
import LoginPage from "../../pages/auth/LoginPage";
import AboutPage from "../../pages/dashboard/AboutPage";
import DeliveryPage from "../../pages/dashboard/DeliveryPage";
import HomePage from "../../pages/dashboard/HomePage";
import { RequireAuth, RequireUnauth } from "./guards";

export const AppRoutes = () => {
  return (
    <BrowserRouter>
      <Suspense fallback={<LoadingPage />}>
        <Routes>
          <Route element={<RequireUnauth />}>
            {/* <Route path="/register" element={<RegisterPage />} /> */}
            <Route path="/login" element={<LoginPage />} />
          </Route>

          <Route element={<RequireAuth />}>
            <Route path="/" element={<HomePage />} />
            <Route path="/about" element={<AboutPage />} />
            <Route path="/analysis/delivery" element={<DeliveryPage />} />
          </Route>

          {/* 404 FALLBACK */}
          {/* <Route
            path="*"
            element={
              <UserLayout>
                <NotFoundPage />
              </UserLayout>
            }
          /> */}
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
};
