-- Companies table
CREATE TABLE companies (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   company_name VARCHAR(255) NOT NULL,
   representative_name VARCHAR(255) NOT NULL,
   phone_number VARCHAR(20) NOT NULL,
   postal_code VARCHAR(8) NOT NULL,
   address TEXT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Users table
CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   company_id UUID NOT NULL REFERENCES companies(id),
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   password_hash VARCHAR(255) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Business partners table
CREATE TABLE business_partners (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   company_id UUID NOT NULL REFERENCES companies(id),
   company_name VARCHAR(255) NOT NULL,
   representative_name VARCHAR(255) NOT NULL,
   phone_number VARCHAR(20) NOT NULL,
   postal_code VARCHAR(8) NOT NULL,
   address TEXT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Bank accounts table
CREATE TABLE bank_accounts (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   business_partner_id UUID NOT NULL REFERENCES business_partners(id),
   bank_name VARCHAR(255) NOT NULL,
   branch_name VARCHAR(255) NOT NULL,
   account_number VARCHAR(20) NOT NULL,
   account_name VARCHAR(255) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Invoices table
CREATE TABLE invoices (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  company_id UUID NOT NULL REFERENCES companies(id),
  business_partner_id UUID NOT NULL REFERENCES business_partners(id),
  issue_date DATE NOT NULL,
  payment_amount INTEGER NOT NULL,
  fee INTEGER NOT NULL,
  fee_rate DECIMAL(5,2) NOT NULL,
  consumption_tax INTEGER NOT NULL,
  tax_rate DECIMAL(5,2) NOT NULL,
  total_amount INTEGER NOT NULL,
  payment_due_date DATE NOT NULL,
  status VARCHAR(20) NOT NULL CHECK (status IN ('unpaid', 'processing', 'paid', 'error')),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- API keys table
CREATE TABLE api_keys (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  api_key VARCHAR(64) NOT NULL UNIQUE,
  expires_at TIMESTAMP WITH TIME ZONE,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
