import React, { useState } from 'react';
import { 
  LayoutDashboard, 
  TrendingUp, 
  History, 
  Wallet, 
  Bell, 
  Search, 
  ArrowUpRight,
  Plus,
  ArrowRight
} from 'lucide-react';
import { motion } from 'framer-motion';

const App: React.FC = () => {
  const [activeTab, setActiveTab] = useState('dashboard');

  const stats = [
    { label: 'Total Investment', value: '₹ 4,50,000', change: '+12.4%' },
    { label: 'Current Value', value: '₹ 5,12,340', change: '+14.2%' },
    { label: 'Total Returns', value: '₹ 62,340', change: '+8.1%' }
  ];

  const recentOrders = [
    { id: '1', fund: 'Quant Small Cap Fund', amount: '₹ 5,000', type: 'SIP', status: 'Completed', date: '2 hours ago' },
    { id: '2', fund: 'Parag Parikh Flexi Cap', amount: '₹ 10,000', type: 'Lumpsum', status: 'Pending', date: '5 hours ago' },
    { id: '3', fund: 'Mirae Asset Large Cap', amount: '₹ 2,500', type: 'SIP', status: 'Completed', date: 'Yesterday' }
  ];

  return (
    <div className="dashboard-container">
      {/* Sidebar */}
      <aside className="sidebar">
        <div className="logo">
          <TrendingUp size={28} />
          <span>FinFlow</span>
        </div>

        <nav className="nav-links">
          {[
            { id: 'dashboard', label: 'Dashboard', icon: LayoutDashboard },
            { id: 'market', label: 'Market', icon: TrendingUp },
            { id: 'portfolio', label: 'Portfolio', icon: Wallet },
            { id: 'orders', label: 'Orders', icon: History },
          ].map((item) => (
            <div 
              key={item.id}
              className={`nav-item ${activeTab === item.id ? 'active' : ''}`}
              onClick={() => setActiveTab(item.id)}
            >
              <item.icon size={20} />
              <span>{item.label}</span>
            </div>
          ))}
        </nav>

        <div style={{ marginTop: 'auto' }}>
          <div className="nav-item">
            <Bell size={20} />
            <span>Notifications</span>
          </div>
        </div>
      </aside>

      {/* Main Content */}
      <main className="main-content">
        <header className="header">
          <div>
            <h1>Welcome back, Investor</h1>
            <p style={{ color: 'var(--text-secondary)', marginTop: '0.5rem' }}>
              Your portfolio is up by 2.4% today.
            </p>
          </div>

          <div style={{ display: 'flex', gap: '1.5rem', alignItems: 'center' }}>
            <div style={{ position: 'relative' }}>
              <Search 
                size={18} 
                style={{ position: 'absolute', left: '12px', top: '50%', transform: 'translateY(-50%)', color: 'var(--text-secondary)' }} 
              />
              <input 
                type="text" 
                placeholder="Search funds..." 
                style={{
                  background: 'var(--card-bg)',
                  border: '1px solid var(--border-color)',
                  padding: '0.6rem 1rem 0.6rem 2.5rem',
                  borderRadius: '12px',
                  color: 'white',
                  width: '240px'
                }}
              />
            </div>
            <div className="user-profile">
              <div className="avatar"></div>
              <span>John Doe</span>
            </div>
          </div>
        </header>

        {/* Stats Grid */}
        <section className="stats-grid">
          {stats.map((stat, i) => (
            <motion.div 
              key={i}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: i * 0.1 }}
              className="stat-card"
            >
              <div className="stat-label">{stat.label}</div>
              <div className="stat-value">{stat.value}</div>
              <div className="stat-change">
                <ArrowUpRight size={16} />
                {stat.change}
              </div>
            </motion.div>
          ))}
        </section>

        {/* Orders Section */}
        <section>
          <div className="section-title">
            <span>Recent Orders</span>
            <button className="btn-primary">
              <Plus size={18} style={{ marginRight: '0.5rem' }} />
              New SIP
            </button>
          </div>

          <table className="orders-table">
            <thead>
              <tr>
                <th>Fund Name</th>
                <th>Amount</th>
                <th>Type</th>
                <th>Status</th>
                <th>Date</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {recentOrders.map((order, i) => (
                <motion.tr 
                  key={order.id}
                  initial={{ opacity: 0 }}
                  animate={{ opacity: 1 }}
                  transition={{ delay: 0.3 + (i * 0.1) }}
                >
                  <td style={{ fontWeight: 500 }}>{order.fund}</td>
                  <td>{order.amount}</td>
                  <td>{order.type}</td>
                  <td>
                    <span className={`status-badge status-${order.status.toLowerCase()}`}>
                      {order.status}
                    </span>
                  </td>
                  <td style={{ color: 'var(--text-secondary)' }}>{order.date}</td>
                  <td>
                    <ArrowRight size={18} style={{ cursor: 'pointer', color: 'var(--text-secondary)' }} />
                  </td>
                </motion.tr>
              ))}
            </tbody>
          </table>
        </section>
      </main>
    </div>
  );
};

export default App;
