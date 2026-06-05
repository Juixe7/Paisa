import React from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';

export default function BudgetScreen() {
  const categories = [
    { name: 'Groceries', actual: 1200, limit: 5000, color: '#10B981' },
    { name: 'Dining Out', actual: 950, limit: 3000, color: '#38BDF8' },
    { name: 'Commute', actual: 600, limit: 2000, color: '#F59E0B' },
    { name: 'Rent & Renters', actual: 15000, limit: 15000, color: '#6366F1' },
    { name: 'Subscriptions', actual: 450, limit: 1000, color: '#EC4899' },
  ];

  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <Text style={styles.headerTitle}>Monthly Budgets</Text>
      <Text style={styles.headerSubtitle}>Track your spending limits vs actual expenses</Text>

      <View style={styles.card}>
        {categories.map((cat, index) => {
          const ratio = Math.min(cat.actual / cat.limit, 1);
          const percent = (ratio * 100).toFixed(0);
          
          return (
            <View key={index} style={styles.budgetRow}>
              <View style={styles.rowHeader}>
                <Text style={styles.catName}>{cat.name}</Text>
                <Text style={styles.catDetails}>
                  ₹{cat.actual} of ₹{cat.limit} ({percent}%)
                </Text>
              </View>

              <View style={styles.barBg}>
                <View style={[styles.barFill, { width: `${percent}%`, backgroundColor: cat.color }]} />
              </View>
            </View>
          );
        })}
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
  },
  headerTitle: {
    color: '#FFF',
    fontSize: 24,
    fontWeight: 'bold',
    marginTop: 8,
  },
  headerSubtitle: {
    color: '#94A3B8',
    fontSize: 14,
    marginTop: 4,
    marginBottom: 24,
  },
  card: {
    backgroundColor: '#1E293B',
    borderRadius: 16,
    padding: 20,
    borderWidth: 1,
    borderColor: '#334155',
  },
  budgetRow: {
    marginBottom: 20,
  },
  rowHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 8,
  },
  catName: {
    color: '#FFF',
    fontWeight: '600',
    fontSize: 15,
  },
  catDetails: {
    color: '#94A3B8',
    fontSize: 13,
  },
  barBg: {
    height: 10,
    backgroundColor: '#0F172A',
    borderRadius: 5,
    overflow: 'hidden',
  },
  barFill: {
    height: '100%',
    borderRadius: 5,
  },
});
