import { useState } from "react";
import { DashboardLayout } from "../components/layout/Dashboard";

// This is the dashboard overview page
// / it will contain the dashboard overview

export const DashboardOverview = () => {
    const [activeTab, setActiveTab] = useState('dashboard');

    return (

        <DashboardLayout>
            
        </DashboardLayout>
    );
};