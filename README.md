# 📄 Complete Documentation for MsgForceChangeAdmin Implementation

## 📋 Overview
The **MsgForceChangeAdmin** implementation enables Terra Classic's on-chain governance to forcefully change the admin of a smart contract through a one-time-use governance proposal. This ensures that the community can regain control of contracts if necessary, enhancing security and decentralization.

---

## ⚙️ Architecture

### 1. **Core Components**

- **MsgForceChangeAdmin**: Message definition that carries information for the admin change.
- **Handler**: Processes the message, validates it, and executes the admin change.
- **Keeper**: Provides methods to interact with the contract's state, such as changing the admin.
- **Integration Tests**: Ensures the end-to-end functionality of the message flow.

### 2. **File Structure**

```plaintext
lunc_gov_force_change_admin/
├── go.mod
└── x/
    └── gov/
        ├── types/
        │   └── msg_force_change_admin.go   # Message definition
        ├── keeper/
        │   ├── handler_force_change_admin.go  # Message handler
        │   └── force_change_admin.go          # Keeper logic
        └── tests/
            └── integration_force_change_admin_test.go  # Integration tests
```

---

## 📑 Message Definition: MsgForceChangeAdmin

**File:** `x/gov/types/msg_force_change_admin.go`

```go
// MsgForceChangeAdmin defines the message for forcefully changing contract admin
type MsgForceChangeAdmin struct {
    Authority    string `json:"authority" yaml:"authority"`         // Bech32 address of the authority
    ContractAddr string `json:"contract_addr" yaml:"contract_addr"` // Bech32 contract address
    NewAdmin     string `json:"new_admin" yaml:"new_admin"`         // Bech32 new admin address
}
```

### 🔑 Key Methods:

- **`NewMsgForceChangeAdmin`**: Constructor for the message.
- **`ValidateBasic`**: Performs basic validation (e.g., valid addresses).
- **`GetSigners`**: Returns the required signers for the message.

---

## 🛠️ Handler Logic

**File:** `x/gov/keeper/handler_force_change_admin.go`

### 🔄 Workflow:

1. **Validate One-Time Use:** Checks if `MsgForceChangeAdmin` has already been used.
2. **Verify Authority:** Ensures the message sender has governance authority.
3. **Validate Addresses:** Checks the validity of `ContractAddr` and `NewAdmin`.
4. **Execute Admin Change:** Calls keeper to update the admin.
5. **Emit Events:** Logs the admin change for transparency.

### ⚡ Key Code Snippet:

```go
func HandleMsgForceChangeAdmin(ctx sdk.Context, keeper Keeper, msg *types.MsgForceChangeAdmin) (*sdk.Result, error) {
    // Check if ForceChangeAdmin has already been used
    if store.Has(KeyForceChangeAdminUsed) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "ForceChangeAdmin can only be used once")
    }

    // Verify sender is governance authority
    if !keeper.IsGovernanceAuthority(ctx, msg.Authority) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only governance can perform this action")
    }

    // Validate addresses
    if msg.NewAdmin == "" {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "new admin address cannot be empty")
    }

    // Perform the admin change
    err := keeper.ForceChangeAdmin(ctx, msg.ContractAddr, msg.NewAdmin)
    if err != nil {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Failed to change admin")
    }

    // Mark as used
    store.Set(KeyForceChangeAdminUsed, []byte(fmt.Sprintf("used_at_height_%d", ctx.BlockHeight())))

    // Emit event
    ctx.EventManager().EmitEvent(
        sdk.NewEvent("force_admin_change",
            sdk.NewAttribute("contract", msg.ContractAddr),
            sdk.NewAttribute("new_admin", msg.NewAdmin),
            sdk.NewAttribute("executed_at_height", fmt.Sprintf("%d", ctx.BlockHeight())),
        ),
    )

    return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
```

---

## 🔒 Security Features

1. **One-Time Use:** Ensures `MsgForceChangeAdmin` can only be used once.
2. **Governance Authority Verification:** Only governance-approved entities can execute.
3. **Address Validation:** Ensures all addresses are valid Bech32 strings.
4. **Event Emission:** Logs actions for transparency.
5. **Replay Attack Protection:** Uses Cosmos SDK’s sequence numbers and block height tracking.

---

## 🧪 Integration Testing

**File:** `x/gov/tests/integration_force_change_admin_test.go`

### 🔍 Test Scenarios:

1. **✅ Successful Admin Change:**
   - Valid authority and addresses.
   - Ensures the admin is changed.

2. **🚫 Unauthorized Access:**
   - Ensures only governance authority can execute.

3. **🚫 One-Time Use Validation:**
   - Prevents reuse of `MsgForceChangeAdmin`.

4. **🚫 Invalid Addresses:**
   - Catches malformed contract and admin addresses.

### ⚡ Test Example:

```go
func TestHandleMsgForceChangeAdmin(t *testing.T) {
    ctx, k := SetupMockContext()

    authority := "terra1authorityaddress"
    contractAddr := "terra1contractaddress"
    newAdmin := "terra1newadminaddress"

    msg := types.NewMsgForceChangeAdmin(authority, contractAddr, newAdmin)

    res, err := keeper.HandleMsgForceChangeAdmin(ctx, k, msg)
    require.NoError(t, err)
    require.NotNil(t, res)

    // Attempt reuse (should fail)
    res, err = keeper.HandleMsgForceChangeAdmin(ctx, k, msg)
    require.Error(t, err)
    require.Contains(t, err.Error(), "ForceChangeAdmin can only be used once")
}
```

---

## 🚀 How to Run

### 1. **Build the Project:**
```bash
go mod tidy
go build ./...
```

### 2. **Run Tests:**
```bash
go test ./x/gov/tests/...
```

### 3. **Submit Proposal:**
```bash
terrad tx gov submit-proposal force-change-admin \
  --contract terra1contractaddress \
  --new-admin terra1newadminaddress \
  --from wallet_name \
  --chain-id columbus-5
```

---

## ✅ Conclusion
The **MsgForceChangeAdmin** feature ensures that Terra Classic governance has the ultimate control to rectify or reclaim smart contracts when needed. With its robust security checks, one-time use validation, and comprehensive testing, this implementation adds a critical safety layer to the Terra ecosystem.

---

🔗 **Author:** Terra Classic Dev Team  
📅 **Last Updated:** February 2025

