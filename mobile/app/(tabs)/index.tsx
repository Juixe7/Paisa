import React from 'react';
import { View, Text, StyleSheet, ScrollView, TouchableOpacity } from 'react-native';

export default function DashboardScreen() {
  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      {/* Balance Section */}
      <View style={styles.balanceCard}>
        <Text style={styles.balanceLabel}>Available Balance This Month</Text>
        <Text style={styles.balanceAmount}>₹10,500.00</Text>
        
        <View style={styles.velocityRow}>
          <View style={[styles.indicator, { backgroundColor: '#10B981' }]} />
          <Text style={styles.velocityText}>Spending Velocity: <Text style={styles.boldText}>On Track</Text> (25 days left)</Text>
        </View>
      </View>

      {/* AI Insight Card */}
      <View style={styles.insightCard}>
        <Text style={styles.insightHeader}>💡 AI Insight</Text>
        <Text style={styles.insightText}>
          You spent 12% less on Swiggy compared to last week. On track to save ₹3,000 this month.
        </Text>
      </View>

      {/* Top Categories */}
      <Text style={styles.sectionTitle}>Top Spending Categories</Text>
      
      <View style={styles.categoryCard}>
        <View style={styles.categoryInfo}>
          <Text style={styles.categoryName}>Groceries</Text>
          <Text style={styles.categorySpend}>₹1,200 / ₹5,000</Text>
        </View>
        <View style={styles.progressBarBg}>
          <View style={[styles.progressBarFill, { width: '24%', backgroundColor: '#10B981' }]} />
        </View>

        <View style={[styles.categoryInfo, { marginTop: 16 }]}>
          <Text style={styles.categoryName}>Dining Out</Text>
          <Text style={styles.categorySpend}>₹950 / ₹3,000</Text>
        </View>
        <View style={styles.progressBarBg}>
          <View style={[styles.progressBarFill, { width: '31.6%', backgroundColor: '#38BDF8' }]} />
        </View>

        <View style={[styles.categoryInfo, { marginTop: 16 }]}>
          <Text style={styles.categoryName}>Commute</Text>
          <Text style={styles.categorySpend}>₹600 / ₹2,000</Text>
        </View>
        <View style={styles.progressBarBg}>
          <View style={[styles.progressBarFill, { width: '30%', backgroundColor: '#F59E0B' }]} />
        </View>
      </View>

      {/* Quick Actions */}
      <View style={styles.actionsRow}>
        <TouchableOpacity style={styles.actionBtn}>
          <Text style={styles.actionBtnText}>+ Add Expense</Text>
        </TouchableOpacity>
        <TouchableOpacity style={[styles.actionBtn, { backgroundColor: '#38BDF8' }]}>
          <Text style={[styles.actionBtnText, { color: '#0F172A' }]}>Pay via UPI</Text>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#0F172A',
  },
  content: {
    padding: 16,
    paddingBottom: 32,
  },
  balanceCard: {
    backgroundColor: '#1E293B',
    padding: 24,
    borderRadius: 16,
    borderWidth: 1,
    borderColor: '#334155',
    marginBottom: 16,
  },
  balanceLabel: {
    color: '#94A3B8',
    fontSize: 14,
    fontWeight: '500',
  },
  balanceAmount: {
    color: '#FFF',
    fontSize: 36,
    fontWeight: 'bold',
    marginTop: 8,
  },
  velocityRow: {
    flexDirection: 'row',
    alignItems: 'center',
    marginTop: 16,
    paddingTop: 16,
    borderTopWidth: 1,
    borderTopColor: '#334155',
  },
  indicator: {
    width: 10,
    height: 10,
    borderRadius: 5,
    marginRight: 8,
  },
  velocityText: {
    color: '#E2E8F0',
    fontSize: 13,
  },
  boldText: {
    fontWeight: 'bold',
    color: '#10B981',
  },
  insightCard: {
    backgroundColor: '#1E1B4B', // Deep indigo
    padding: 16,
    borderRadius: 12,
    borderWidth: 1,
    borderColor: '#312E81',
    marginBottom: 24,
  },
  insightHeader: {
    color: '#A5B4FC',
    fontWeight: 'bold',
    fontSize: 14,
    marginBottom: 6,
  },
  insightText: {
    color: '#E0E7FF',
    fontSize: 14,
    lineHeight: 20,
  },
  sectionTitle: {
    color: '#FFF',
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 12,
  },
  categoryCard: {
    backgroundColor: '#1E293B',
    padding: 20,
    borderRadius: 16,
    borderWidth: 1,
    borderColor: '#334155',
    marginBottom: 24,
  },
  categoryInfo: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: 8,
  },
  categoryName: {
    color: '#E2E8F0',
    fontWeight: '600',
    fontSize: 14,
  },
  categorySpend: {
    color: '#94A3B8',
    fontSize: 13,
  },
  progressBarBg: {
    height: 8,
    backgroundColor: '#0F172A',
    borderRadius: 4,
    overflow: 'hidden',
  },
  progressBarFill: {
    height: '100%',
    borderRadius: 4,
  },
  actionsRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    gap: 12,
  },
  actionBtn: {
    flex: 1,
    backgroundColor: '#1E293B',
    paddingVertical: 14,
    borderRadius: 10,
    alignItems: 'center',
    borderWidth: 1,
    borderColor: '#334155',
  },
  actionBtnText: {
    color: '#FFF',
    fontWeight: 'bold',
    fontSize: 15,
  },
});
