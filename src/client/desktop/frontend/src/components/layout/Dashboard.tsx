import { Sidebar, SidebarHeader, SidebarContent, SidebarGroup } from "../ui/sidebar";
import { Button } from "../ui/button";
import { useState } from "react";

// This is the base layout for the dashboard
// / it will contain the sidebar and the main content area
// / the sidebar will contain the navigation links
// / the main content area will contain the page content
const sidebarTabs = [
    { id: 'dashboard', label: 'Dashboard' },
    { id: 'campaigns', label: 'Campaigns' },
    { id: 'stagers', label: 'Stagers' },
    { id: 'agents', label: 'Agents' },
    { id: 'listeners', label: 'Listeners' },
    { id: 'credentials', label: 'Credentials' },
    { id: 'settings', label: 'Settings' }
];

export const DashboardLayout = () => {
    const [activeTab, setActiveTab] = useState('dashboard');

    return (
        <Sidebar>
            <SidebarHeader>
                <div className="flex-1">
                    <h1>Dashboard</h1>
                </div>
            </SidebarHeader>
            <SidebarGroup>
                <SidebarContent>
                    <div className="flex flex-col h-full">
                        {sidebarTabs.map(tab => (
                            <Button
                                key={tab.id}
                                variant={activeTab === tab.id ? 'default' : 'ghost'}
                                className="w-full justify-start"
                                onClick={() => setActiveTab(tab.id)}
                            >
                                {tab.label}
                            </Button>
                        ))}
                    </div>
                </SidebarContent>
            </SidebarGroup>
        </Sidebar>
    );
};