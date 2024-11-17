# **LeetGo**

LeetGo is a CLI tool designed to help you explore tagged questions from LeetCode, grouped by companies. It provides features to list companies, view problems associated with a company, and apply filters like frequency, difficulty, and acceptance rate.

---

## **Usage**

### **Basic Commands**
1. **List All Companies:**
   ```bash
    $ leetgo -l
    Available Companies:
    accenture
    accolite
    activision
    adobe
   ```

2. **List Problems for a Specific Company:**
   ```bash
    $ leetgo -c google
    Problems for company 'google':
    https://leetcode.com/problems/race-car | Frequency: 100.00%
    https://leetcode.com/problems/stock-price-fluctuation | Frequency: 91.87%
    https://leetcode.com/problems/my-calendar-i | Frequency: 90.55%
   ```
   Displays problems tagged with the specified company. Add the `-a` flag for detailed output.

3. **Apply Filters:**
   Combine flags like `--frequency`, `--difficulty`, `--acceptance`, and `--non-premium` for refined results.

---

## **Download and Installation**

Pre-built binaries are available for **Windows**, **Linux**, and **macOS** in the [Releases](https://github.com/yuisofull/leetgo/releases) section.

### **Steps to Install**
If you'd prefer to build the binary yourself, follow these steps:
```bash
git clone https://github.com/yuisofull/leetgo.git
cd leetgo
go build -o leetgo
```

---

## **Supported Platforms**

| Platform   | File Name       |
|------------|-----------------|
| **Windows**| `leetgo.exe`    |
| **Linux**  | `leetgo`        |
| **macOS**  | `leetgo-darwin` |

Download the appropriate file for your system from the [Releases](https://github.com/yuisofull/leetgo/releases) section.

---

### **List Problems for a Specific Company**

```bash
leetgo -c <company-name>
```
**Description:** Lists all problems tagged with the given company name.

**Example (Brief Output):**
```bash
$ leetgo -c -f google
Problems for company 'google':
https://leetcode.com/problems/race-car | Frequency: 100.00%
https://leetcode.com/problems/stock-price-fluctuation | Frequency: 91.87%
https://leetcode.com/problems/my-calendar-i | Frequency: 90.55%
```

**Example (Detailed Output):**
```bash
$ leetgo -c google -a
Problems for company 'google':
ID: 1
Title: Two Sum
URL: https://leetcode.com/problems/two-sum/
Difficulty: Easy
Frequency: 95.00%
Acceptance: 45.67%
Premium: false

ID: 2
Title: Add Two Numbers
URL: https://leetcode.com/problems/add-two-numbers/
Difficulty: Medium
Frequency: 89.50%
Acceptance: 38.24%
Premium: true
```

---

### **Filter Problems**

#### **Sort by Frequency**
```bash
leetgo -c <company-name> --frequency
```
Lists problems sorted by frequency.

**Example:**
```bash
leetgo -c google --frequency
```

#### **Filter by Difficulty**
```bash
leetgo -c <company-name> --difficulty <easy|medium|hard>
```
Lists problems filtered by difficulty level.

**Example:**
```bash
leetgo -c google --difficulty medium
```

#### **Sort by Acceptance Rate**
```bash
leetgo -c <company-name> --acceptance
```
Lists problems sorted by acceptance rate.

**Example:**
```bash
leetgo -c google --acceptance
```

#### **Exclude Premium Problems**
```bash
leetgo -c <company-name> --non-premium
```
Lists problems excluding premium ones.

**Example:**
```bash
leetgo -c google --non-premium
```

#### **Limit Number of Problems**
```bash
leetgo -c <company-name> --limit <number>
```
Limits the number of problems returned.

**Example:**
```bash
leetgo -c google --limit 3
```

#### **Combine Filters**
Filters can be combined for more precise results.

**Example:**
```bash
leetgo -c google --frequency --non-premium --limit 5 -a
```

---

### **Help Command**
For a detailed list of all commands and flags, use:
```bash
leetgo --help
```

## **Error Handling**
- Invalid company name:
  ```bash
  leetgo -c invalid-company
  ```
  **Output:**
  ```
  Error fetching problems for company 'invalid-company'
  ```
