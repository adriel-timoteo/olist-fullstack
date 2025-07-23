import { Suspense } from "react";
import { BrowserRouter, Route, Routes } from "react-router";
import LoadingPage from "../../components/LoadingPage";
import HomePage from "../../pages/dashboard/HomePage";
import AboutPage from "../../pages/dashboard/AboutPage";
import AnalysisPage from "../../pages/dashboard/AnalysisPage";

export const AppRoutes = () => {
  return (
    <BrowserRouter>
      <Suspense fallback={<LoadingPage />}>
        <Routes>
          {/* <Route element={<RequireUnauth />}> */}
          {/* <Route path="/register" element={<RegisterPage />} />
          <Route path="/login" element={<LoginPage />} /> */}
          {/* </Route> */}

          {/* <Route element={<RequireAuth />}> */}
          <Route path="/" element={<HomePage />} />
          <Route path="/about" element={<AboutPage />} />
          <Route path="/analysis" element={<AnalysisPage />} />

          {/* </Route> */}
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
